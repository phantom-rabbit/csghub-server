package workflows

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"go.temporal.io/sdk/activity"
	"gopkg.in/yaml.v2"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/git/gitserver/gitea"
	"opencsg.com/csghub-server/builder/parquet"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/builder/store/s3"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
	dvCom "opencsg.com/csghub-server/dataviewer/common"
)

type DataViewerActivity interface {
	BeginViewerJob(ctx context.Context) error
	GetCardFromReadme(ctx context.Context, req types.UpdateViewerReq) (*dvCom.CardData, error)
	ScanRepoFiles(ctx context.Context, scanParam dvCom.ScanRepoFileReq) (*dvCom.RepoFilesClass, error)
	DetermineCardData(ctx context.Context, determineParam dvCom.DetermineCardReq) (*dvCom.CardData, error)
	CheckIfNeedRebuild(ctx context.Context, checkParam dvCom.CheckBuildReq) (bool, error)
	CreateParquetBranch(ctx context.Context, req types.UpdateViewerReq) (string, error)
	CopyParquetFiles(ctx context.Context, copyReq dvCom.CopyParquetReq) (*dvCom.CardData, error)
	DownloadSplitFiles(ctx context.Context, downloadReq dvCom.DownloadFileReq) (*dvCom.DownloadCard, error)
	ConvertToParquetFiles(ctx context.Context, convertReq dvCom.ConvertReq) error
	UploadParquetFiles(ctx context.Context, uploadReq dvCom.UploadParquetReq) (*dvCom.CardData, error)
	UpdateCardData(ctx context.Context, cardReq dvCom.UpdateCardReq) error
	CleanUp(ctx context.Context, req types.UpdateViewerReq) error
	UpdateWorkflowStatus(ctx context.Context, status dvCom.UpdateWorkflowStatusReq) error
}

type dataViewerActivityImpl struct {
	gitServer   gitserver.GitServer
	s3Client    s3.Client
	cfg         *config.Config
	viewerStore database.DataviewerStore
}

func NewDataViewerActivity(cfg *config.Config, gs gitserver.GitServer) (DataViewerActivity, error) {
	s3Client, err := s3.NewMinio(cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to init s3 client for data viewer, error: %w", err)
	}
	return &dataViewerActivityImpl{
		gitServer:   gs,
		s3Client:    s3Client,
		cfg:         cfg,
		viewerStore: database.NewDataviewerStore(),
	}, nil
}

func (dva *dataViewerActivityImpl) BeginViewerJob(ctx context.Context) error {
	wfCtx := activity.GetInfo(ctx)
	workflowID := wfCtx.WorkflowExecution.ID
	runID := wfCtx.WorkflowExecution.RunID
	job, err := dva.viewerStore.GetJob(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("get viewer job by workflowID %s for beginning, cause: %w", workflowID, err)
	}

	job.RunID = runID
	job.Status = types.WorkflowRunning
	job.Logs = types.WorkflowMsgRunning
	job.StartTime = time.Now()

	_, err = dva.viewerStore.UpdateJob(ctx, *job)
	if err != nil {
		slog.Error("update viewer job info for beginning", slog.Any("job", job), slog.Any("error", err))
		return fmt.Errorf("update viewer job info by workflowID %s for beginning, cause: %w", workflowID, err)
	}
	return nil
}

func (dva *dataViewerActivityImpl) GetCardFromReadme(ctx context.Context, req types.UpdateViewerReq) (*dvCom.CardData, error) {
	fileReq := gitserver.GetRepoInfoByPathReq{
		Namespace: req.Namespace,
		Name:      req.Name,
		Ref:       req.Branch,
		Path:      types.REPOCARD_FILENAME,
		RepoType:  req.RepoType,
	}
	f, err := dva.gitServer.GetRepoFileContents(context.Background(), fileReq)
	if err != nil {
		return nil, fmt.Errorf("get %s repo %s/%s branch %s file %s content error: %w", fileReq.RepoType, fileReq.Namespace, fileReq.Name, fileReq.Ref, fileReq.Path, err)
	}
	slog.Debug("getRepoCardData", slog.Any("f.Content", f.Content))
	decodedContent, err := base64.StdEncoding.DecodeString(f.Content)
	if err != nil {
		return nil, fmt.Errorf("decode %s repo %s/%s branch %s file %s content, error: %w", fileReq.RepoType, fileReq.Namespace, fileReq.Name, fileReq.Ref, fileReq.Path, err)
	}
	matches := dvCom.REG.FindStringSubmatch(string(decodedContent))
	yamlString := ""
	if len(matches) > 1 {
		yamlString = matches[1]
	} else {
		return nil, fmt.Errorf("%s repo %s/%s branch %s card yaml config is empty due to invalid content", fileReq.RepoType, fileReq.Namespace, fileReq.Name, fileReq.Ref)
	}

	var card dvCom.CardData
	err = yaml.Unmarshal([]byte(yamlString), &card)
	if err != nil {
		return nil, fmt.Errorf("unmarshal %s repo %s/%s branch %s yaml error: %w, decoded content: %s", fileReq.RepoType, fileReq.Namespace, fileReq.Name, fileReq.Ref, err, yamlString)
	}
	return &card, nil
}

func (dva *dataViewerActivityImpl) ScanRepoFiles(ctx context.Context, scanParam dvCom.ScanRepoFileReq) (*dvCom.RepoFilesClass, error) {
	repoReq := dvCom.RepoFilesReq{
		Namespace: scanParam.Req.Namespace,
		RepoName:  scanParam.Req.Name,
		RepoType:  scanParam.Req.RepoType,
		Ref:       scanParam.Req.Branch,
		Folder:    "",
		GSTree:    dva.gitServer.GetRepoFileTree,
	}
	fileClass := dvCom.RepoFilesClass{
		AllFiles:     make(map[string]*types.File),
		ParquetFiles: make(map[string]*types.File),
		JsonlFiles:   make(map[string]*types.File),
		CsvFiles:     make(map[string]*types.File),
	}
	err := GetFilePaths(repoReq, &fileClass, scanParam.MaxFileSize)
	if err != nil {
		return nil, fmt.Errorf("scan repo file error: %w", err)
	}
	return &fileClass, nil
}

func (dva *dataViewerActivityImpl) autoBuildCardData(card *dvCom.CardData, sortKeys []string, targetFiles map[string]*types.File) {
	var (
		trainFiles []string
		testFiles  []string
		valFiles   []string
		otherFiles []string

		trainFileObjects []dvCom.FileObject
		testFileObjects  []dvCom.FileObject
		valFileObjects   []dvCom.FileObject
		otherFileObjects []dvCom.FileObject
	)
	for _, path := range sortKeys {
		file, exists := targetFiles[path]
		if !exists {
			continue
		}
		if IsTrainFile(path) {
			trainFiles = append(trainFiles, path)
			trainFileObjects = append(trainFileObjects, TransferFileObject(file, DefaultSubsetName, TrainSplitName))
		} else if IsTestFile(path) {
			testFiles = append(testFiles, path)
			testFileObjects = append(testFileObjects, TransferFileObject(file, DefaultSubsetName, TestSplitName))
		} else if IsValidationFile(path) {
			valFiles = append(valFiles, path)
			valFileObjects = append(valFileObjects, TransferFileObject(file, DefaultSubsetName, ValSplitName))
		} else {
			otherFiles = append(otherFiles, path)
			otherFileObjects = append(otherFileObjects, TransferFileObject(file, DefaultSubsetName, OtherSplitName))
		}
	}
	var configData dvCom.ConfigData
	var datasetInfo dvCom.DatasetInfo
	if len(trainFiles) > 0 {
		configData.DataFiles = append(configData.DataFiles,
			dvCom.DataFiles{Split: TrainSplitName, Path: trainFiles},
		)
		datasetInfo.Splits = append(datasetInfo.Splits,
			dvCom.Split{Name: TrainSplitName, NumExamples: 0, Origins: trainFileObjects},
		)
	}
	if len(testFiles) > 0 {
		configData.DataFiles = append(configData.DataFiles,
			dvCom.DataFiles{Split: TestSplitName, Path: testFiles},
		)
		datasetInfo.Splits = append(datasetInfo.Splits,
			dvCom.Split{Name: TestSplitName, NumExamples: 0, Origins: testFileObjects},
		)
	}
	if len(valFiles) > 0 {
		configData.DataFiles = append(configData.DataFiles,
			dvCom.DataFiles{Split: ValSplitName, Path: valFiles},
		)
		datasetInfo.Splits = append(datasetInfo.Splits,
			dvCom.Split{Name: ValSplitName, NumExamples: 0, Origins: valFileObjects},
		)
	}
	if len(otherFiles) > 0 {
		configData.DataFiles = append(configData.DataFiles,
			dvCom.DataFiles{Split: OtherSplitName, Path: otherFiles},
		)
		datasetInfo.Splits = append(datasetInfo.Splits,
			dvCom.Split{Name: OtherSplitName, NumExamples: 0, Origins: otherFileObjects},
		)
	}
	if len(configData.DataFiles) > 0 {
		configData.ConfigName = DefaultSubsetName
		datasetInfo.ConfigName = DefaultSubsetName
		card.Configs = append(card.Configs, configData)
		card.DatasetInfos = append(card.DatasetInfos, datasetInfo)
	}
}

func (dva *dataViewerActivityImpl) fillUpCardData(card *dvCom.CardData, sortKeys []string, targetFiles map[string]*types.File) *dvCom.CardData {
	var configs []dvCom.ConfigData
	var infos []dvCom.DatasetInfo
	for _, conf := range card.Configs {
		var datafiles []dvCom.DataFiles
		var splits []dvCom.Split
		for _, datafile := range conf.DataFiles {
			var newPath interface{}
			reqFiles := GetPatternFileList(datafile.Path)
			if len(reqFiles) > 0 {
				newPath = reqFiles
			} else {
				slog.Warn("datafile.Path is not either string or []interface{})", slog.Any("datafile.Path", datafile.Path))
				newPath = datafile.Path
			}
			datafiles = append(datafiles, dvCom.DataFiles{Split: datafile.Split, Path: newPath})
			realReqFiles := ConvertRealFiles(reqFiles, sortKeys, targetFiles, conf.ConfigName, datafile.Split)
			splits = append(splits, dvCom.Split{Name: datafile.Split, NumExamples: 0, Origins: realReqFiles})
		}
		configs = append(configs, dvCom.ConfigData{ConfigName: conf.ConfigName, DataFiles: datafiles})
		infos = append(infos, dvCom.DatasetInfo{ConfigName: conf.ConfigName, Splits: splits})
	}
	return &dvCom.CardData{Configs: configs, DatasetInfos: infos}
}

func (dva *dataViewerActivityImpl) DetermineCardData(ctx context.Context, determineParam dvCom.DetermineCardReq) (*dvCom.CardData, error) {
	var scopeFiles map[string]*types.File
	if determineParam.RepoDataType == RepoParquetData {
		scopeFiles = determineParam.Class.ParquetFiles
	} else if determineParam.RepoDataType == RepoJsonData {
		scopeFiles = determineParam.Class.JsonlFiles
	} else if determineParam.RepoDataType == RepoCsvData {
		scopeFiles = determineParam.Class.CsvFiles
	}
	if scopeFiles == nil || len(scopeFiles) < 1 {
		slog.Warn("no target valid files found", slog.Any("card", determineParam.Card))
		return nil, nil
	}

	keys := make([]string, 0, len(scopeFiles))
	for key := range scopeFiles {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if determineParam.Card.Configs == nil {
		dva.autoBuildCardData(&determineParam.Card, keys, scopeFiles)
		return &determineParam.Card, nil
	} else {
		newCard := dva.fillUpCardData(&determineParam.Card, keys, scopeFiles)
		return newCard, nil
	}
}

func (dva *dataViewerActivityImpl) CheckIfNeedRebuild(ctx context.Context, checkParam dvCom.CheckBuildReq) (bool, error) {
	if checkParam.Card.Configs == nil {
		slog.Warn("card configs is nil, no need to rebuild", slog.Any("req", checkParam.Req), slog.Any("card", checkParam.Card))
		return false, nil
	}
	viewer, err := dva.viewerStore.GetViewerByRepoID(ctx, checkParam.Req.RepoID)
	if err != nil {
		slog.Error("get viewer for compare card", slog.Any("req", checkParam.Req),
			slog.Any("repo id", checkParam.Req.RepoID), slog.Any("error", err))
		return true, nil
	}
	if viewer == nil {
		return true, nil
	}
	newMD5 := GetCardDataMD5(checkParam.Card)
	if viewer.DataviewerJob.CardMD5 == newMD5 {
		slog.Warn("card data md5 not changed, no need to rebuild", slog.Any("req", checkParam.Req),
			slog.Any("card", checkParam.Card), slog.Any("newMD5", newMD5))
		return false, nil
	}
	return true, nil
}

func (dva *dataViewerActivityImpl) CreateParquetBranch(ctx context.Context, req types.UpdateViewerReq) (string, error) {
	findReq := gitserver.GetBranchReq{
		Namespace: req.Namespace,
		Name:      req.Name,
		Ref:       dvCom.ParquetBranch,
		RepoType:  req.RepoType,
	}
	branch, err := dva.gitServer.GetRepoBranchByName(ctx, findReq)
	if err != nil {
		slog.Warn("get branch by name", slog.Any("branch", findReq.Ref), slog.Any("error", err))
	}
	if err == nil && branch != nil {
		deleteReq := gitserver.DeleteBranchReq{
			Namespace: req.Namespace,
			Name:      req.Name,
			Ref:       findReq.Ref,
			RepoType:  req.RepoType,
			Username:  GitDefaultUserName,
			Email:     GitDefaultUserEmail,
		}
		err = dva.gitServer.DeleteRepoBranch(ctx, deleteReq)
		if err != nil {
			slog.Error("failed to delete branch", slog.Any("req", req), slog.Any("req", req), slog.Any("err", err))
			return branch.Name, fmt.Errorf("delete branch %s error: %w", findReq.Ref, err)
		}
	}

	// Create .gitattributes file in new branch
	createReq := &types.CreateFileReq{
		Username:  GitDefaultUserName,
		Email:     GitDefaultUserEmail,
		Message:   "create branch",
		Content:   base64.StdEncoding.EncodeToString([]byte(types.DatasetGitattributesContent)),
		NewBranch: findReq.Ref,
		Namespace: req.Namespace,
		Name:      req.Name,
		FilePath:  types.GitattributesFileName,
		RepoType:  req.RepoType,
	}

	err = dva.gitServer.CreateRepoFile(createReq)
	if err != nil {
		slog.Error("failed to create .gitattributes file in branch", slog.Any("branch", createReq.NewBranch), slog.Any("req", req), slog.Any("error", err))
		return "", fmt.Errorf("failed to create .gitattributes file in branch %s, cause: %w", createReq.NewBranch, err)
	}
	return createReq.NewBranch, nil
}

func (dva *dataViewerActivityImpl) CopyParquetFiles(ctx context.Context, copyReq dvCom.CopyParquetReq) (*dvCom.CardData, error) {
	r, err := parquet.NewS3Reader(ctx, dva.cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create duckdb reader, cause: %w", err)
	}
	cardData := dvCom.CardData{}
	var datasetInfos []dvCom.DatasetInfo
	for _, info := range copyReq.ComputedCardData.DatasetInfos {
		newInfo := dvCom.DatasetInfo{}
		newSplits := []dvCom.Split{}
		for _, split := range info.Splits {
			newSplit := dvCom.Split{}
			newFiles := []dvCom.FileObject{}
			objectNames := []string{}
			count := 0
			if split.Origins != nil {
				for idx, orginfile := range split.Origins {
					getFileContentReq := gitserver.GetRepoInfoByPathReq{
						Namespace: copyReq.Req.Namespace,
						Name:      copyReq.Req.Name,
						Ref:       copyReq.Req.Branch,
						Path:      orginfile.RepoFile,
						RepoType:  copyReq.Req.RepoType,
					}
					f, err := dva.gitServer.GetRepoFileContents(context.Background(), getFileContentReq)
					if err != nil {
						slog.Error("failed to get file content", slog.Any("file", orginfile.RepoFile),
							slog.Any("branch", copyReq.Req.Branch), slog.Any("req", copyReq.Req), slog.Any("error", err))
						return nil, fmt.Errorf("failed to get file %s content in branch %s, cause: %w", orginfile.RepoFile,
							copyReq.Req.Branch, err)
					}
					newFileName := fmt.Sprintf("%s/%s/%05d.parquet", info.ConfigName, split.Name, idx)
					createReq := &types.CreateFileReq{
						Username:  GitDefaultUserName,
						Email:     GitDefaultUserEmail,
						Message:   fmt.Sprintf("submit %s file", newFileName),
						Content:   f.Content,
						Branch:    copyReq.NewBranch,
						Namespace: copyReq.Req.Namespace,
						Name:      copyReq.Req.Name,
						FilePath:  newFileName,
						RepoType:  copyReq.Req.RepoType,
					}
					err = dva.gitServer.CreateRepoFile(createReq)
					if err != nil {
						slog.Error("failed to create file in branch", slog.Any("newfile", newFileName),
							slog.Any("newBranch", copyReq.NewBranch), slog.Any("req", copyReq.Req), slog.Any("error", err))
						return nil, fmt.Errorf("failed to create new file %s in new branch %s, cause: %w",
							newFileName, copyReq.NewBranch, err)
					}
					newFiles = append(newFiles, dvCom.FileObject{
						RepoFile:        newFileName,
						Size:            orginfile.Size,
						Lfs:             orginfile.Lfs,
						LfsRelativePath: orginfile.LfsRelativePath,
					})
					objectNames = append(objectNames, filepath.Join("lfs", orginfile.LfsRelativePath))
				}
				count, err = r.RowCount(ctx, objectNames, types.QueryReq{}, true)
				if err != nil {
					slog.Error("get S3 row count error", slog.Any("req", copyReq.Req),
						slog.Any("config", info.ConfigName), slog.Any("split", split.Name),
						slog.Any("objectNames", objectNames), slog.Any("error", err))
					return nil, fmt.Errorf("failed to get row count for repo %s/%s submit %s split %s, cause: %w",
						copyReq.Req.Namespace, copyReq.Req.Name, info.ConfigName, split.Name, err)
				}
			}
			newSplit.Name = split.Name
			newSplit.Files = newFiles
			newSplit.NumExamples = count
			newSplit.Origins = split.Origins
			newSplits = append(newSplits, newSplit)
		}
		newInfo.ConfigName = info.ConfigName
		newInfo.Splits = newSplits
		datasetInfos = append(datasetInfos, newInfo)
	}
	cardData.Configs = copyReq.ComputedCardData.Configs
	cardData.DatasetInfos = datasetInfos
	return &cardData, nil
}

func (dva *dataViewerActivityImpl) DownloadSplitFiles(ctx context.Context, downloadReq dvCom.DownloadFileReq) (*dvCom.DownloadCard, error) {
	var subsets []dvCom.DownloadSubset
	cacheRepoPath := GetCacheRepoPath(ctx, dva.cfg.DataViewer.CacheDir, downloadReq.Req)
	_, err := os.Stat(cacheRepoPath)
	if err != nil && !os.IsNotExist(err) {
		slog.Warn("check cach repo path error for download split file", slog.Any("cacheRepoPath", cacheRepoPath), slog.Any("error", err))
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(cacheRepoPath, os.ModePerm)
		if err != nil {
			slog.Error("create cache repo path error for download split file", slog.Any("cacheRepoPath", cacheRepoPath), slog.Any("error", err))
			return nil, fmt.Errorf("failed to create cache repo path %s for download split file, cause: %w", cacheRepoPath, err)
		}
	} else {
		err = os.RemoveAll(cacheRepoPath)
		if err != nil {
			slog.Error("remove cache repo path error for download split file", slog.Any("cacheRepoPath", cacheRepoPath), slog.Any("error", err))
			return nil, fmt.Errorf("failed to remove cache repo path %s for download split file, cause: %w", cacheRepoPath, err)
		}
		err = os.MkdirAll(cacheRepoPath, os.ModePerm)
		if err != nil {
			slog.Error("create cache repo path error for download split file", slog.Any("cacheRepoPath", cacheRepoPath), slog.Any("error", err))
			return nil, fmt.Errorf("failed to create cache repo path %s for download split file, cause: %w", cacheRepoPath, err)
		}
	}

	for _, info := range downloadReq.Card.DatasetInfos {
		newSplits := []dvCom.DownloadSplit{}
		for _, split := range info.Splits {
			newFiles := []dvCom.FileObject{}
			for idx, file := range split.Origins {
				extName := filepath.Ext(file.RepoFile)
				localFileName := fmt.Sprintf("%05d%s", idx, extName)
				downloadObj, err := dva.downloadFile(ctx, downloadReq.Req, file, &dvCom.FileObject{
					RepoFile:      file.RepoFile,
					LastCommit:    file.LastCommit,
					Lfs:           file.Lfs,
					LocalRepoPath: cacheRepoPath,
					LocalFileName: localFileName,
					Size:          file.Size,
					SubsetName:    info.ConfigName,
					SplitName:     split.Name,
				})
				if err != nil {
					slog.Error("failed to download file", slog.Any("req", downloadReq.Req),
						slog.Any("file", file), slog.Any("error", err))
					return nil, fmt.Errorf("failed to download file %s in branch %s, cause: %w",
						file.RepoFile, downloadReq.Req.Branch, err)
				}
				// slog.Info("downloaded file successfully", slog.Any("req", req), slog.Any("file", file), slog.Any("downloadObj", downloadObj))
				newFiles = append(newFiles, *downloadObj)
			}
			splitPath := fmt.Sprintf("%s/%s/%s", cacheRepoPath, info.ConfigName, split.Name)
			exportPath := fmt.Sprintf("%s_export", splitPath)
			newSplit := dvCom.DownloadSplit{Name: split.Name, LocalPath: splitPath, ExportPath: exportPath, Files: newFiles}
			newSplits = append(newSplits, newSplit)
		}
		newSubset := dvCom.DownloadSubset{ConfigName: info.ConfigName, Splits: newSplits}
		subsets = append(subsets, newSubset)
	}
	return &dvCom.DownloadCard{Configs: downloadReq.Card.Configs, Subsets: subsets}, nil
}

func (dva *dataViewerActivityImpl) downloadFile(ctx context.Context, req types.UpdateViewerReq, orginFile dvCom.FileObject, loadFile *dvCom.FileObject) (*dvCom.FileObject, error) {
	cacheFilePath := fmt.Sprintf("%s/%s/%s", loadFile.LocalRepoPath, loadFile.SubsetName, loadFile.SplitName)
	_, err := os.Stat(cacheFilePath)
	if err != nil && !os.IsNotExist(err) {
		slog.Warn("check cache file path error", slog.Any("cacheFilePath", cacheFilePath), slog.Any("error", err))
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(cacheFilePath, os.ModePerm)
		if err != nil {
			slog.Error("create cache file path error for download file", slog.Any("cacheFilePath", cacheFilePath), slog.Any("error", err))
			return nil, fmt.Errorf("failed to create cache file path %s for download file, cause: %w", cacheFilePath, err)
		}
	}
	localFileFullPath := fmt.Sprintf("%s/%s", cacheFilePath, loadFile.LocalFileName)
	if orginFile.Lfs {
		err := dva.downloadLfsFile(ctx, localFileFullPath, orginFile, loadFile)
		if err != nil {
			return nil, fmt.Errorf("fail to download repo %s/%s lfs file %s, error: %w", req.Namespace, req.Name, orginFile.RepoFile, err)
		}
	} else {
		err := dva.downloadNormalFile(ctx, localFileFullPath, req, orginFile, loadFile)
		if err != nil {
			return nil, fmt.Errorf("fail to download repo %s/%s normal file %s, error: %w", req.Namespace, req.Name, orginFile.RepoFile, err)
		}
	}

	return loadFile, nil
}

func (dva *dataViewerActivityImpl) downloadLfsFile(ctx context.Context, localFileFullPath string, orginFile dvCom.FileObject, loadFile *dvCom.FileObject) error {
	objectKey := orginFile.LfsRelativePath
	objectKey = path.Join("lfs", objectKey)
	loadFile.ObjectKey = objectKey

	if !dva.cfg.DataViewer.DownloadLfsFile {
		slog.Warn("skip download lfs file", slog.Any("file", orginFile))
		return nil
	}

	reqParams := make(url.Values)
	signedUrl, err := dva.s3Client.PresignedGetObject(ctx, dva.cfg.S3.Bucket, objectKey, types.OssFileExpire, reqParams)
	if err != nil {
		return fmt.Errorf("fail to get lfs file download url error: %w", err)
	}
	resp, err := http.Get(signedUrl.String())
	if err != nil {
		return fmt.Errorf("failed to do http request url %s, error: %w", signedUrl.String(), err)
	}
	defer resp.Body.Close()

	writeFile, err := os.Create(localFileFullPath)
	if err != nil {
		return fmt.Errorf("failed to create local file %s, error: %w", localFileFullPath, err)
	}
	defer writeFile.Close()

	_, err = io.Copy(writeFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save local file %s, error: %w", localFileFullPath, err)
	}
	return nil
}

func (dva *dataViewerActivityImpl) downloadNormalFile(ctx context.Context, localFileFullPath string, req types.UpdateViewerReq, orginFile dvCom.FileObject, loadFile *dvCom.FileObject) error {
	getFileReaderReq := gitserver.GetRepoInfoByPathReq{
		Namespace: req.Namespace,
		Name:      req.Name,
		Ref:       req.Branch,
		Path:      orginFile.RepoFile,
		RepoType:  req.RepoType,
	}
	reader, size, err := dva.gitServer.GetRepoFileReader(ctx, getFileReaderReq)
	loadFile.Size = size
	if err != nil {
		return fmt.Errorf("failed to get repo %s/%s file %s for reader, error: %w", req.Namespace, req.Name, orginFile.RepoFile, err)
	}
	writeFile, err := os.Create(localFileFullPath)
	if err != nil {
		return fmt.Errorf("failed to create local file %s for repo %s/%s file %s, error: %w", localFileFullPath, req.Namespace, req.Name, orginFile.RepoFile, err)
	}
	defer writeFile.Close()

	_, err = io.Copy(writeFile, reader)
	if err != nil {
		return fmt.Errorf("failed to save local file %s for repo %s/%s file %s, error: %w", localFileFullPath, req.Namespace, req.Name, orginFile.RepoFile, err)
	}
	return nil
}

func (dva *dataViewerActivityImpl) ConvertToParquetFiles(ctx context.Context, convertReq dvCom.ConvertReq) error {
	var err error
	writer, err := parquet.NewS3Writer(ctx, dva.cfg)
	if err != nil {
		return fmt.Errorf("failed to create duckdb reader, cause: %w", err)
	}
	for _, subset := range convertReq.DownloadCard.Subsets {
		for _, split := range subset.Splits {
			if len(split.Files) < 1 {
				continue
			}
			objectNames := []string{}
			for _, file := range split.Files {
				if file.Lfs && !dva.cfg.DataViewer.DownloadLfsFile {
					objectNames = append(objectNames, fmt.Sprintf("'s3://%s/%s'", dva.cfg.S3.Bucket, file.ObjectKey))
				} else {
					objectNames = append(objectNames, fmt.Sprintf("'%s/%s/%s/%s'", file.LocalRepoPath, file.SubsetName, file.SplitName, file.LocalFileName))
				}
			}
			slog.Debug("ConvertToParquetFiles", slog.Any("objectNames", objectNames))
			_, err = os.Stat(split.ExportPath)
			if err != nil && !os.IsNotExist(err) {
				slog.Warn("check export file path error", slog.Any("ExportPath", split.ExportPath), slog.Any("error", err))
			}
			if os.IsNotExist(err) {
				err = os.MkdirAll(split.ExportPath, os.ModePerm)
				if err != nil {
					return fmt.Errorf("failed to create export path %s for convert, error: %w", split.ExportPath, err)
				}
			}
			method := ""
			if convertReq.RepoDataType == RepoJsonData {
				method = "read_json_auto"
			} else if convertReq.RepoDataType == RepoCsvData {
				method = "read_csv_auto"
			}
			splitExportPath := fmt.Sprintf("%s/", strings.TrimSuffix(split.ExportPath, "/"))
			err = writer.ConvertToParquet(ctx, method, objectNames, dva.cfg.DataViewer.ThreadNumOfExport, splitExportPath)
			if err != nil {
				slog.Error("failed to convert data", slog.Any("objectNames", objectNames),
					slog.Any("req", convertReq.Req), slog.Any("configname", subset.ConfigName),
					slog.Any("split", split), slog.Any("error", err))
				return fmt.Errorf("failed to convert data for repo %s/%s, subset %s, split: %v, cause: %w",
					convertReq.Req.Namespace, convertReq.Req.Name, subset.ConfigName, split, err)
			}
			slog.Debug("convert parquet success", slog.Any("req", convertReq.Req),
				slog.Any("subset", subset.ConfigName), slog.Any("split", split.Name),
				slog.Any("exportPath", split.ExportPath))
		}
	}
	return nil
}

func (dva *dataViewerActivityImpl) UploadParquetFiles(ctx context.Context, uploadReq dvCom.UploadParquetReq) (*dvCom.CardData, error) {
	r, err := parquet.NewS3Reader(ctx, dva.cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create duckdb reader, cause: %w", err)
	}
	cardData := dvCom.CardData{}
	var datasetInfos []dvCom.DatasetInfo
	for _, subset := range uploadReq.DownloadCard.Subsets {
		newInfo := dvCom.DatasetInfo{}
		newSplits := []dvCom.Split{}
		for _, split := range subset.Splits {
			newSplit := dvCom.Split{}
			newFiles := []dvCom.FileObject{}
			if split.ExportPath == "" {
				continue
			}
			entries, err := os.ReadDir(split.ExportPath)
			if err != nil {
				slog.Error("read export dir error", slog.Any("exportPath", split.ExportPath), slog.Any("error", err))
				return nil, fmt.Errorf("failed to read dir %s, cause: %w", split.ExportPath, err)
			}
			sort.Slice(entries, func(i, j int) bool {
				nameI := entries[i].Name()
				nameJ := entries[j].Name()
				numI, _ := strconv.Atoi(nameI[:len(nameI)-len(filepath.Ext(nameI))])
				numJ, _ := strconv.Atoi(nameJ[:len(nameJ)-len(filepath.Ext(nameJ))])
				return numI < numJ
			})
			objectNames := []string{}
			for _, entry := range entries {
				fileName := entry.Name()
				if entry.Type().IsRegular() && IsValidParquetFile(fileName) {
					extName := filepath.Ext(fileName)
					fileNameSeq := fileName[:len(fileName)-len(extName)]
					fileNameSeqInt, err := strconv.Atoi(fileNameSeq)
					if err != nil {
						slog.Warn("invalid file name to int error", slog.Any("fileName", fileName), slog.Any("ExportPath", split.ExportPath), slog.Any("error", err))
						continue
					}
					fileInfo, err := entry.Info()
					if err != nil {
						slog.Error("get file info error", slog.Any("filePath", fileName), slog.Any("error", err))
						return nil, fmt.Errorf("failed to get file info %s under path %s, cause: %w", fileName, split.ExportPath, err)
					}
					realFile := fmt.Sprintf("%s/%s", split.ExportPath, fileName)
					newFileName := fmt.Sprintf("%05d%s", fileNameSeqInt, extName)
					uploadFile := &dvCom.FileObject{
						ConvertPath: realFile,
						RepoFile:    fmt.Sprintf("%s/%s/%s", subset.ConfigName, split.Name, newFileName),
						Size:        fileInfo.Size(),
						Lfs:         true,
					}
					err = dva.uploadToRepo(ctx, uploadReq.Req, uploadFile, uploadReq.NewBranch)
					if err != nil {
						slog.Error("upload file to repo error", slog.Any("realFile", realFile),
							slog.Any("req", uploadReq.Req), slog.Any("newbranch", uploadReq.NewBranch), slog.Any("error", err))
						return nil, fmt.Errorf("failed to upload file %s to repo %s/%s branch %s, cause: %w", realFile,
							uploadReq.Req.Namespace, uploadReq.Req.Name, uploadReq.NewBranch, err)
					}
					newFiles = append(newFiles, *uploadFile)
					objectNames = append(objectNames, realFile)
				}
			}
			count, err := r.RowCount(ctx, objectNames, types.QueryReq{}, false)
			if err != nil {
				slog.Error("get row count error", slog.Any("req", uploadReq.Req),
					slog.Any("config", subset.ConfigName), slog.Any("split", split.Name),
					slog.Any("objectNames", objectNames), slog.Any("error", err))
				return nil, fmt.Errorf("failed to get row count for repo %s/%s submit %s split %s, cause: %w",
					uploadReq.Req.Namespace, uploadReq.Req.Name, subset.ConfigName, split.Name, err)
			}
			newSplit.Name = split.Name
			newSplit.Files = newFiles
			newSplit.NumExamples = count
			newSplit.Origins = split.Files
			newSplits = append(newSplits, newSplit)
		}
		newInfo.ConfigName = subset.ConfigName
		newInfo.Splits = newSplits
		datasetInfos = append(datasetInfos, newInfo)
	}
	cardData.Configs = uploadReq.DownloadCard.Configs
	cardData.DatasetInfos = datasetInfos
	return &cardData, nil
}

func (dva *dataViewerActivityImpl) uploadToRepo(ctx context.Context, req types.UpdateViewerReq, uploadFile *dvCom.FileObject, newBranch string) error {
	f, err := os.Open(uploadFile.ConvertPath)
	if err != nil {
		return fmt.Errorf("open file %s, cause: %w", uploadFile.ConvertPath, err)
	}
	defer f.Close()

	pointer, err := gitea.GeneratePointer(f)
	if err != nil {
		return fmt.Errorf("fail to get lfs file %s point, cause: %w", uploadFile.ConvertPath, err)
	}
	encodingLfsContent := base64.StdEncoding.EncodeToString([]byte(pointer.StringContent()))

	_, err = f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seek to beginning of file %s, cause: %w", uploadFile.ConvertPath, err)
	}

	uploadFile.LfsRelativePath = pointer.RelativePath()
	objectKey := filepath.Join("lfs", pointer.RelativePath())
	uploadInfo, err := dva.s3Client.PutObject(ctx, dva.cfg.S3.Bucket, objectKey, f, pointer.Size, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("upload file %s to S3, cause: %w", uploadFile.ConvertPath, err)
	}

	if uploadInfo.Size != pointer.Size {
		return fmt.Errorf("uploaded S3 file %s size does not match expected size: %d != %d", uploadFile.ConvertPath, uploadInfo.Size, pointer.Size)
	}

	createReq := &types.CreateFileReq{
		Username:  GitDefaultUserName,
		Email:     GitDefaultUserEmail,
		Message:   "upload parquet file",
		FilePath:  uploadFile.RepoFile,
		Content:   encodingLfsContent,
		Namespace: req.Namespace,
		Name:      req.Name,
		RepoType:  req.RepoType,
		Branch:    newBranch,
	}

	err = dva.gitServer.CreateRepoFile(createReq)
	if err != nil {
		return fmt.Errorf("failed to create lfs file %s in repo %s/%s branch %s, cause: %w", uploadFile.RepoFile, req.Namespace, req.Name, newBranch, err)
	}
	return nil
}

func (dva *dataViewerActivityImpl) UpdateCardData(ctx context.Context, cardReq dvCom.UpdateCardReq) error {
	wfCtx := activity.GetInfo(ctx)
	workflowID := wfCtx.WorkflowExecution.ID

	job, err := dva.viewerStore.GetJob(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("get job by workflow id %s for update card, cause: %w", workflowID, err)
	}

	finalCardDataJson, err := json.Marshal(cardReq.FinalCardData)
	if err != nil {
		slog.Error("failed to marshal final card data", slog.Any("req", cardReq.Req),
			slog.Any("finalCardData", cardReq.FinalCardData), slog.Any("error", err))
		return fmt.Errorf("marshal final card data to json, cause: %w", err)
	}

	job.AutoCard = (cardReq.OriginCardData.Configs == nil)
	job.CardData = string(finalCardDataJson)
	job.CardMD5 = GetCardDataMD5(cardReq.FinalCardData)

	_, err = dva.viewerStore.UpdateJob(ctx, *job)
	if err != nil {
		return fmt.Errorf("update job by id %d for update card, cause: %w", job.ID, err)
	}
	return nil
}

func (dva *dataViewerActivityImpl) CleanUp(ctx context.Context, req types.UpdateViewerReq) error {
	cacheRepoPath := GetCacheRepoPath(ctx, dva.cfg.DataViewer.CacheDir, req)
	if _, err := os.Stat(cacheRepoPath); !os.IsNotExist(err) {
		err := os.RemoveAll(cacheRepoPath)
		if err != nil {
			slog.Warn("clean up cache repo path %s, error: %w", cacheRepoPath, err)
		}
	}
	return nil
}

func (dva *dataViewerActivityImpl) UpdateWorkflowStatus(ctx context.Context, status dvCom.UpdateWorkflowStatusReq) error {
	wfCtx := activity.GetInfo(ctx)
	workflowID := wfCtx.WorkflowExecution.ID
	runID := wfCtx.WorkflowExecution.RunID

	if status.WorkflowErr != nil {
		slog.Error("run data viewer workflow error", slog.Any("workflowID", workflowID), slog.Any("runID", runID),
			slog.Any("status", status), slog.Any("workflowErr", status.WorkflowErr))
	}

	job, err := dva.viewerStore.GetJob(ctx, workflowID)
	if err != nil {
		slog.Error("get workflow for ending", slog.Any("workflowID", workflowID), slog.Any("err", err))
		return nil
	}

	if status.WorkflowErr != nil {
		job.Status = types.WorkflowFailed
		job.Logs = status.WorkflowErr.Error()
	} else {
		job.Status = types.WorkflowDone
		job.Logs = types.WorkflowMsgDone
	}

	job.EndTime = time.Now()

	_, err = dva.viewerStore.UpdateJob(ctx, *job)
	if err != nil {
		slog.Error("update workflow result for ending", slog.Any("workflowID", workflowID), slog.Any("job", job), slog.Any("error", err))
	}

	if status.WorkflowErr != nil || !status.ShouldUpdateViewer {
		return nil
	}

	viewer, err := dva.viewerStore.GetViewerByRepoID(ctx, status.Req.RepoID)
	if err != nil {
		slog.Error("get viewer workflow for ending", slog.Any("status", status), slog.Any("err", err))
		return nil
	}

	viewer.WorkflowID = workflowID
	_, err = dva.viewerStore.UpdateViewer(ctx, *viewer)
	if err != nil {
		slog.Error("update viewer for workflow ending", slog.Any("status", status), slog.Any("viewer", viewer), slog.Any("error", err))
	}
	return nil
}
