package component

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/parquet"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
	hubCom "opencsg.com/csghub-server/component"
	dvCom "opencsg.com/csghub-server/dataviewer/common"
	"opencsg.com/csghub-server/dataviewer/workflows"
)

type DatasetViewerComponent interface {
	ViewParquetFile(ctx context.Context, req *dvCom.ViewParquetFileReq) (*dvCom.ViewParquetFileResp, error)
	Rows(ctx context.Context, req *dvCom.ViewParquetFileReq, viewerReq types.DataViewerReq) (*dvCom.ViewParquetFileResp, error)
	GetCatalog(ctx context.Context, req *dvCom.ViewParquetFileReq) (*dvCom.CataLogRespone, error)
}

type datasetViewerComponentImpl struct {
	repoStore     database.RepoStore
	repoComponent hubCom.RepoComponent
	gitServer     gitserver.GitServer
	preader       parquet.Reader
	cfg           *config.Config
	viewerStore   database.DataviewerStore
}

func NewDatasetViewerComponent(cfg *config.Config, gs gitserver.GitServer) (DatasetViewerComponent, error) {
	r, err := hubCom.NewRepoComponentImpl(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create repo component,cause:%w", err)
	}

	return &datasetViewerComponentImpl{
		gitServer:     gs,
		cfg:           cfg,
		repoComponent: r,
		repoStore:     database.NewRepoStore(),
		viewerStore:   database.NewDataviewerStore(),
	}, nil
}

func (c *datasetViewerComponentImpl) lazyInit(ctx context.Context) error {
	if c.preader != nil {
		return nil
	}
	r, err := parquet.NewS3Reader(ctx, c.cfg)
	if err != nil {
		c.preader = nil
		return fmt.Errorf("failed to create parquet reader,cause: %w", err)
	}
	c.preader = r
	return nil
}

func (c *datasetViewerComponentImpl) ViewParquetFile(ctx context.Context, req *dvCom.ViewParquetFileReq) (*dvCom.ViewParquetFileResp, error) {
	err := c.lazyInit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init parquet reader, %w", err)
	}

	r, err := c.repoStore.FindByPath(ctx, types.DatasetRepo, req.Namespace, req.RepoName)
	if err != nil {
		return nil, fmt.Errorf("failed to find dataset, error: %w", err)
	}

	allow, err := c.repoComponent.AllowReadAccessRepo(ctx, r, req.CurrentUser)
	if err != nil {
		return nil, fmt.Errorf("failed to check dataset permission, error: %w", err)
	}
	if !allow {
		return nil, hubCom.ErrForbidden
	}
	req.Branch = r.DefaultBranch
	objName, err := c.getParquetObject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("no valid parquet file in request, %w", err)
	}

	sqlReq := types.QueryReq{
		PageSize:  req.Per,
		PageIndex: req.Page,
		Search:    "",
		Where:     "",
		Orderby:   "",
	}

	total, err := c.preader.RowCount(ctx, []string{objName}, sqlReq, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count of parquet file, %w", err)
	}

	columns, columnsType, rows, err := c.preader.FetchRows(ctx, []string{objName}, sqlReq, true)
	if err != nil {
		return nil, fmt.Errorf("failed to view parquet rows, %w", err)
	}

	resp := &dvCom.ViewParquetFileResp{
		Columns:     columns,
		ColumnsType: columnsType,
		Rows:        rows,
		Total:       total,
	}
	return resp, nil
}

func (c *datasetViewerComponentImpl) Rows(ctx context.Context, req *dvCom.ViewParquetFileReq, viewerReq types.DataViewerReq) (*dvCom.ViewParquetFileResp, error) {
	r, err := c.repoStore.FindByPath(ctx, types.DatasetRepo, req.Namespace, req.RepoName)
	if err != nil {
		return nil, fmt.Errorf("failed to find dataset, error: %w", err)
	}

	allow, err := c.repoComponent.AllowReadAccessRepo(ctx, r, req.CurrentUser)
	if err != nil {
		return nil, fmt.Errorf("failed to check dataset permission, error: %w", err)
	}
	if !allow {
		return nil, hubCom.ErrForbidden
	}
	req.Branch = r.DefaultBranch

	err = c.lazyInit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init parquet reader, %w", err)
	}

	_, parquetObjs, err := c.getViewerCardData(ctx, r.ID, &viewerReq)
	if err != nil {
		slog.Warn("do not found viewer card data", slog.Any("repo_id", r.ID), slog.Any("req", req), slog.Any("error", err))
		parquetObjs, err = c.getRepoParquetObjs(ctx, req, viewerReq)
		if err != nil {
			return nil, fmt.Errorf("failed to get parquet objects, %w", err)
		}
	}

	if len(parquetObjs) < 1 {
		return nil, fmt.Errorf("no valid parquet file in request for row data")
	}

	sqlReq := types.QueryReq{
		PageSize:  req.Per,
		PageIndex: req.Page,
		Search:    viewerReq.Search,
		Where:     viewerReq.Where,
		Orderby:   viewerReq.Orderby,
	}

	total, err := c.preader.RowCount(ctx, parquetObjs, sqlReq, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count of parquet file, %w", err)
	}
	columns, columnsType, rows, err := c.preader.FetchRows(ctx, parquetObjs, sqlReq, true)
	if err != nil {
		return nil, fmt.Errorf("failed to view parquet rows, %w", err)
	}

	resp := &dvCom.ViewParquetFileResp{
		Columns:     columns,
		ColumnsType: columnsType,
		Rows:        rows,
		Total:       total,
		Orderby:     viewerReq.Orderby,
		Where:       viewerReq.Where,
		Search:      viewerReq.Search,
	}
	return resp, nil
}

func (c *datasetViewerComponentImpl) getRepoParquetObjs(ctx context.Context, req *dvCom.ViewParquetFileReq, viewerReq types.DataViewerReq) ([]string, error) {
	cardData, err := c.getDatasetCatalog(ctx, req, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get dataset catalog, error: %w", err)
	}
	slog.Debug("Rows cardData", slog.Any("cardData", cardData))
	var reqFiles []string
	var hasWildCard bool
	var tree []string

	for _, config := range cardData.Configs {
		if config.ConfigName == viewerReq.Config {
			for _, datafile := range config.DataFiles {
				if datafile.Split == viewerReq.Split {
					reqFiles, hasWildCard = c.getPatternFileList(datafile.Path)
					break
				}
			}
			break
		}
	}

	realReqFiles := reqFiles
	if hasWildCard {
		// need get real files match test/test-* in repo
		realReqFiles, _ = c.convertRealFiles(req, reqFiles, tree)
	}
	parquetObjs := c.getFilesOBJs(ctx, req, realReqFiles)
	if len(parquetObjs) < 1 {
		return nil, fmt.Errorf("no valid files in request")
	}
	return parquetObjs, nil
}

func (c *datasetViewerComponentImpl) GetCatalog(ctx context.Context, req *dvCom.ViewParquetFileReq) (*dvCom.CataLogRespone, error) {
	r, err := c.repoStore.FindByPath(ctx, types.DatasetRepo, req.Namespace, req.RepoName)
	if err != nil {
		return nil, fmt.Errorf("failed to find dataset, error: %w", err)
	}

	allow, err := c.repoComponent.AllowReadAccessRepo(ctx, r, req.CurrentUser)
	if err != nil {
		return nil, fmt.Errorf("failed to check dataset permission, error: %w", err)
	}
	if !allow {
		return nil, hubCom.ErrForbidden
	}
	req.Branch = r.DefaultBranch

	err = c.lazyInit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init parquet reader, %w", err)
	}

	viewerCardData, _, err := c.getViewerCardData(ctx, r.ID, nil)
	if err != nil {
		slog.Warn("do not found viewer card data", slog.Any("repo_id", r.ID), slog.Any("req", req), slog.Any("error", err))
	} else {
		return viewerCardData, nil
	}

	cardData, err := c.getDatasetCatalog(ctx, req, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get dataset catalog, error: %w", err)
	}
	return &dvCom.CataLogRespone{
		Configs:      cardData.Configs,
		DatasetInfos: cardData.DatasetInfos,
	}, nil
}

func (c *datasetViewerComponentImpl) getViewerCardData(ctx context.Context, repoID int64, viewerReq *types.DataViewerReq) (*dvCom.CataLogRespone, []string, error) {
	viewer, err := c.viewerStore.GetViewerByRepoID(ctx, repoID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get viewer by repo_id %d, error: %w", repoID, err)
	}

	if viewer == nil || viewer.DataviewerJob == nil || len(viewer.DataviewerJob.CardData) < 1 {
		return nil, nil, fmt.Errorf("viewer card data is empty")
	}

	if viewer.DataviewerJob.Status == types.WorkflowPending || viewer.DataviewerJob.Status == types.WorkflowFailed {
		return &dvCom.CataLogRespone{
			Status: viewer.DataviewerJob.Status,
			Logs:   viewer.DataviewerJob.Logs,
		}, nil, nil
	}

	var viewerCardData dvCom.CardData
	err = json.Unmarshal([]byte(viewer.DataviewerJob.CardData), &viewerCardData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal viewer card data, error: %w", err)
	}

	parquetObjs := []string{}
	var newInfos []dvCom.DatasetInfo

	for _, info := range viewerCardData.DatasetInfos {
		newSplits := []dvCom.Split{}
		for _, split := range info.Splits {
			if viewerReq != nil && viewerReq.Config == info.ConfigName && viewerReq.Split == split.Name {
				for _, f := range split.Files {
					parquetObjs = append(parquetObjs, fmt.Sprintf("lfs/%s", f.LfsRelativePath))
				}
			}
			newSplit := dvCom.Split{
				Name:        split.Name,
				NumExamples: split.NumExamples,
				Files:       nil,
				Origins:     nil,
			}
			newSplits = append(newSplits, newSplit)
		}
		info.Splits = newSplits
		newInfos = append(newInfos, info)
	}
	viewerCardData.DatasetInfos = newInfos

	return &dvCom.CataLogRespone{
		Configs:      viewerCardData.Configs,
		DatasetInfos: viewerCardData.DatasetInfos,
		Status:       viewer.DataviewerJob.Status,
		Logs:         viewer.DataviewerJob.Logs,
	}, parquetObjs, nil
}

func (c *datasetViewerComponentImpl) getDatasetCatalog(ctx context.Context, req *dvCom.ViewParquetFileReq, calcTotal bool) (*dvCom.CardData, error) {
	cardData, err := c.getRepoCardData(ctx, req, calcTotal)
	if err == nil {
		return cardData, nil
	}
	slog.Warn("cannot get card data for repo", slog.Any("req", req), slog.Any("error", err))
	cardData, err = c.autoGenerateCatalog(ctx, req, calcTotal)
	if err != nil {
		return nil, fmt.Errorf("failed to auto gen catalog, error: %w", err)
	}
	return cardData, nil
}

func (c *datasetViewerComponentImpl) getRepoCardData(ctx context.Context, req *dvCom.ViewParquetFileReq, calcTotal bool) (*dvCom.CardData, error) {
	getFileContentReq := gitserver.GetRepoInfoByPathReq{
		Namespace: req.Namespace,
		Name:      req.RepoName,
		Ref:       req.Branch,
		Path:      types.REPOCARD_FILENAME,
		RepoType:  types.DatasetRepo,
	}
	f, err := c.gitServer.GetRepoFileContents(ctx, getFileContentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get readme.md contents, cause:%w", err)
	}
	slog.Debug("getRepoCardData", slog.Any("f.Content", f.Content))
	decodedContent, err := base64.StdEncoding.DecodeString(f.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode readme.md contents, cause:%w", err)
	}
	decodedContentStr := string(decodedContent)
	matches := dvCom.REG.FindStringSubmatch(decodedContentStr)
	yamlString := ""
	if len(matches) > 1 {
		yamlString = matches[1]
	} else {
		return nil, fmt.Errorf("repo card yaml configs is empty")
	}

	var card dvCom.CardData
	err = yaml.Unmarshal([]byte(yamlString), &card)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal readme.md yaml contents, cause: %w, decodedContent: %v", err, yamlString)
	}
	slog.Debug("Unmarshal", slog.Any("card", card))
	if card.Configs == nil {
		return nil, fmt.Errorf("repo card data configs is empty")
	}
	if calcTotal && card.DatasetInfos == nil {
		card = c.generateCardDatasetInfo(ctx, req, card)
	}
	return &card, nil
}

func (c *datasetViewerComponentImpl) generateCardDatasetInfo(ctx context.Context, req *dvCom.ViewParquetFileReq, card dvCom.CardData) dvCom.CardData {
	var configs []dvCom.ConfigData
	var infos []dvCom.DatasetInfo
	var tree []string
	for _, conf := range card.Configs {
		var datafiles []dvCom.DataFiles
		var splits []dvCom.Split
		for _, datafile := range conf.DataFiles {
			var newPath interface{}
			reqFiles, hasWildCard := c.getPatternFileList(datafile.Path)
			if len(reqFiles) > 0 {
				newPath = reqFiles
			} else {
				slog.Warn("datafile.Path is not either string or []interface{})", slog.Any("datafile.Path", datafile.Path))
				newPath = datafile.Path
			}
			datafiles = append(datafiles, dvCom.DataFiles{Split: datafile.Split, Path: newPath})
			realReqFiles := reqFiles
			if hasWildCard {
				realReqFiles, tree = c.convertRealFiles(req, reqFiles, tree)
			}
			total := c.getFilesRowCount(ctx, req, realReqFiles)
			splits = append(splits, dvCom.Split{Name: datafile.Split, NumExamples: total})
		}
		configs = append(configs, dvCom.ConfigData{ConfigName: conf.ConfigName, DataFiles: datafiles})
		infos = append(infos, dvCom.DatasetInfo{ConfigName: conf.ConfigName, Splits: splits})
	}
	return dvCom.CardData{Configs: configs, DatasetInfos: infos}
}

func (c *datasetViewerComponentImpl) getPatternFileList(path interface{}) ([]string, bool) {
	if path == nil {
		return []string{}, false
	}
	var (
		files       []string
		hasWildCard bool
	)
	hasWildCard = false
	if slice, ok := path.([]interface{}); ok {
		for _, v := range slice {
			files = append(files, v.(string))
			if strings.HasSuffix(v.(string), dvCom.WILDCARD) {
				hasWildCard = true
			}
		}
	} else if slice, ok := path.([]string); ok {
		files = slice
		for _, v := range slice {
			if strings.HasSuffix(v, dvCom.WILDCARD) {
				hasWildCard = true
				break
			}
		}
	} else {
		files = []string{path.(string)}
		if strings.HasSuffix(path.(string), dvCom.WILDCARD) {
			hasWildCard = true
		}
	}
	return files, hasWildCard
}

func (c *datasetViewerComponentImpl) convertRealFiles(req *dvCom.ViewParquetFileReq, splitFiles []string, tree []string) ([]string, []string) {
	var err error
	if len(tree) < 1 {
		// skip get all tree
		tree, err = hubCom.GetFilePaths(req.Namespace, req.RepoName, "", types.DatasetRepo, req.Branch, c.gitServer.GetRepoFileTree)
		if err != nil {
			slog.Error("Failed to get repo file paths", slog.Any("req", req), slog.Any("error", err))
			return splitFiles, tree
		}
		if tree == nil {
			return splitFiles, tree
		}
	}

	var phyFiles []string

	for _, filePattern := range splitFiles {
		if !strings.Contains(filePattern, dvCom.WILDCARD) {
			phyFiles = append(phyFiles, filePattern)
		} else {
			fileReg, err := regexp.Compile(filePattern)
			if err != nil {
				slog.Warn("invalid regexp format", slog.Any("filePattern", filePattern), slog.Any("err", err))
				phyFiles = append(phyFiles, filePattern)
				continue
			}
			for _, repoFile := range tree {
				// repo file match like: test/test-* and end with .parquet
				if fileReg.MatchString(repoFile) && workflows.IsValidParquetFile(repoFile) {
					phyFiles = append(phyFiles, repoFile)
				}
			}
		}
	}
	slog.Debug("convertRealFiles", slog.Any("splitFiles", splitFiles), slog.Any("phyFiles", phyFiles))
	return phyFiles, tree
}

func (c *datasetViewerComponentImpl) autoGenerateCatalog(ctx context.Context, req *dvCom.ViewParquetFileReq, calcTotal bool) (*dvCom.CardData, error) {
	tree, err := hubCom.GetFilePaths(req.Namespace, req.RepoName, "", types.DatasetRepo, req.Branch, c.gitServer.GetRepoFileTree)
	if err != nil {
		return nil, fmt.Errorf("failed to get repo tree, error: %w", err)
	}
	if tree == nil {
		return nil, fmt.Errorf("failed to find any files")
	}
	slog.Debug("get tree", slog.Any("tree", tree))
	cardData := c.genDefaultCatalog(ctx, req, tree, calcTotal)
	return &cardData, nil
}

func (c *datasetViewerComponentImpl) genDefaultCatalog(ctx context.Context, req *dvCom.ViewParquetFileReq, tree []string, calcTotal bool) dvCom.CardData {
	var (
		trainFiles []string
		testFiles  []string
		valFiles   []string
	)
	var configData dvCom.ConfigData
	var datasetInfo dvCom.DatasetInfo
	for _, item := range tree {
		if !workflows.IsValidParquetFile(item) {
			continue
		}
		if workflows.IsTrainFile(strings.ToLower(item)) {
			trainFiles = append(trainFiles, item)
		} else if workflows.IsTestFile(strings.ToLower(item)) {
			testFiles = append(testFiles, item)
		} else if workflows.IsValidationFile(strings.ToLower(item)) {
			valFiles = append(valFiles, item)
		}
	}
	if len(trainFiles) > 0 {
		total := 0
		if calcTotal {
			total = c.getFilesRowCount(ctx, req, trainFiles)
		}
		configData.DataFiles = append(configData.DataFiles, dvCom.DataFiles{Split: workflows.TrainSplitName, Path: trainFiles})
		datasetInfo.Splits = append(datasetInfo.Splits, dvCom.Split{Name: workflows.TrainSplitName, NumExamples: total})
	}
	if len(testFiles) > 0 {
		total := 0
		if calcTotal {
			total = c.getFilesRowCount(ctx, req, testFiles)
		}
		configData.DataFiles = append(configData.DataFiles, dvCom.DataFiles{Split: workflows.TestSplitName, Path: testFiles})
		datasetInfo.Splits = append(datasetInfo.Splits, dvCom.Split{Name: workflows.TestSplitName, NumExamples: total})
	}
	if len(valFiles) > 0 {
		total := 0
		if calcTotal {
			total = c.getFilesRowCount(ctx, req, valFiles)
		}
		configData.DataFiles = append(configData.DataFiles, dvCom.DataFiles{Split: workflows.ValSplitName, Path: valFiles})
		datasetInfo.Splits = append(datasetInfo.Splits, dvCom.Split{Name: workflows.ValSplitName, NumExamples: total})
	}
	configData.ConfigName = workflows.DefaultSubsetName
	datasetInfo.ConfigName = workflows.DefaultSubsetName
	var configList []dvCom.ConfigData
	var dsInfoList []dvCom.DatasetInfo
	if len(configData.DataFiles) > 0 {
		configList = append(configList, configData)
		dsInfoList = append(dsInfoList, datasetInfo)
	}
	return dvCom.CardData{Configs: configList, DatasetInfos: dsInfoList}
}

func (c *datasetViewerComponentImpl) getFilesRowCount(ctx context.Context, req *dvCom.ViewParquetFileReq, files []string) int {
	slog.Debug("getFilesRowCount", slog.Any("files", files))
	parquetObjs := c.getFilesOBJs(ctx, req, files)
	if len(parquetObjs) < 1 {
		return 0
	}
	sqlReq := types.QueryReq{
		PageSize:  10,
		PageIndex: 1,
		Search:    "",
		Orderby:   "",
	}

	total, err := c.preader.RowCount(ctx, parquetObjs, sqlReq, true)
	if err != nil {
		slog.Warn("failed to get parquet row counts", slog.Any("parquetObjs", parquetObjs), slog.Any("error", err))
	}
	return total
}

func (c *datasetViewerComponentImpl) getFilesOBJs(ctx context.Context, req *dvCom.ViewParquetFileReq, files []string) []string {
	var parquetObjs []string
	for _, file := range files {
		objName, err := c.getParquetObject(ctx, &dvCom.ViewParquetFileReq{
			Namespace: req.Namespace,
			RepoName:  req.RepoName,
			Branch:    req.Branch,
			Path:      file,
		})
		if err != nil {
			slog.Warn("failed to get parquet object name", slog.Any("req", req), slog.Any("error", err))
			continue
		}
		parquetObjs = append(parquetObjs, objName)
	}
	return parquetObjs
}

func (c *datasetViewerComponentImpl) getParquetObject(ctx context.Context, req *dvCom.ViewParquetFileReq) (string, error) {
	getFileContentReq := gitserver.GetRepoInfoByPathReq{
		Namespace: req.Namespace,
		Name:      req.RepoName,
		Ref:       req.Branch,
		Path:      req.Path,
		RepoType:  types.DatasetRepo,
	}
	f, err := c.gitServer.GetRepoFileContents(ctx, getFileContentReq)
	if err != nil {
		return "", fmt.Errorf("failed to get file contents,cause:%v", err)
	}
	if f.LfsRelativePath == "" {
		return "", fmt.Errorf("file LfsRelativePath is empty for %s", getFileContentReq.Path)
	}
	return "lfs/" + f.LfsRelativePath, nil
}
