package handler

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/common/utils/common"
	dvCom "opencsg.com/csghub-server/dataviewer/common"
	"opencsg.com/csghub-server/dataviewer/component"
)

type DatasetViewerHandler struct {
	c component.DatasetViewerComponent
}

func NewDatasetViewerHandler(cfg *config.Config, gs gitserver.GitServer) (*DatasetViewerHandler, error) {
	dvc, err := component.NewDatasetViewerComponent(cfg, gs)
	if err != nil {
		return nil, err
	}

	return &DatasetViewerHandler{
		c: dvc,
	}, nil
}

// GetDatasetDemoData godoc
// @Security     ApiKey
// @Summary      Get the demo data of the dataset
// @Description  get the demo data of the dataset
// @Tags         Dataset
// @Accept       json
// @Produce      json
// @Param        namespace path string true "namespace"
// @Parsm        name path string true "name"
// @Param        file_path path string true "file_path"
// @Param        per query int false "per" default(50)
// @Param        page query int false "per page" default(1)
// @Success      200  {object}  types.Response{} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /datasets/{namespace}/{name}/viewer/{file_path} [get]
func (h *DatasetViewerHandler) View(ctx *gin.Context) {
	currentUser := httpbase.GetCurrentUser(ctx)
	namespace, name, err := common.GetNamespaceAndNameFromContext(ctx)
	if err != nil {
		slog.Error("Bad repo request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		slog.Error("Bad paging request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	req := new(dvCom.ViewParquetFileReq)
	req.Namespace = namespace
	req.RepoName = name
	req.Path = ctx.Param("file_path")
	req.Per = per
	req.Page = page
	req.CurrentUser = currentUser

	resp, err := h.c.ViewParquetFile(ctx.Request.Context(), req)
	if err != nil {
		slog.Error("Failed to view parquet file", "error", err)
		httpbase.ServerError(ctx, err)
		return
	}

	httpbase.OK(ctx, resp)
}

// GetDatasetCatalog godoc
// @Security     ApiKey
// @Summary      Get catalog of the dataset
// @Description  get catalog of the dataset
// @Tags         Dataset
// @Accept       json
// @Produce      json
// @Param        namespace path string true "namespace"
// @Parsm        name path string true "name"
// @Success      200  {object}  types.Response{} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /datasets/{namespace}/{name}/dataviewer/catalog [get]
func (h *DatasetViewerHandler) Catalog(ctx *gin.Context) {
	currentUser := httpbase.GetCurrentUser(ctx)
	namespace, name, err := common.GetNamespaceAndNameFromContext(ctx)
	if err != nil {
		slog.Error("Bad request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	req := new(dvCom.ViewParquetFileReq)
	req.Namespace = namespace
	req.RepoName = name
	req.CurrentUser = currentUser

	catalog, err := h.c.GetCatalog(ctx.Request.Context(), req)
	if err != nil {
		slog.Error("Failed to get dataset catalog", slog.Any("req", req), slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, catalog)
}

// GetDatasetCatalog godoc
// @Security     ApiKey
// @Summary      Get catalog of the dataset
// @Description  get catalog of the dataset
// @Tags         Dataset
// @Accept       json
// @Produce      json
// @Param        namespace path string true "namespace"
// @Parsm        name path string true "name"
// @Param        config query string true "config"
// @Param        split query string true "split"
// @Param        search query string false "search"
// @Param        where query string false "where"
// @Param        orderby query string false "orderby"
// @Param        per query int false "per" default(50)
// @Param        page query int false "per page" default(1)
// @Success      200  {object}  types.Response{} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /datasets/{namespace}/{name}/dataviewer/rows [get]
func (h *DatasetViewerHandler) Rows(ctx *gin.Context) {
	currentUser := httpbase.GetCurrentUser(ctx)
	namespace, name, err := common.GetNamespaceAndNameFromContext(ctx)
	if err != nil {
		slog.Error("Bad repo request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}

	per, page, err := common.GetPerAndPageFromContext(ctx)
	if err != nil {
		slog.Error("Bad paging request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}

	// config=SelfRC&split=train
	config := ctx.Query("config")
	split := ctx.Query("split")
	search := ctx.Query("search")
	where := ctx.Query("where")
	orderby := ctx.Query("orderby")
	if config == "" || split == "" {
		slog.Error("Bad view rows request format")
		httpbase.BadRequest(ctx, "Bad view rows request format")
		return
	}
	var viewReq types.DataViewerReq
	viewReq.Config = config
	viewReq.Split = split
	viewReq.Search = search
	viewReq.Where = where
	viewReq.Orderby = orderby

	req := new(dvCom.ViewParquetFileReq)
	req.Namespace = namespace
	req.RepoName = name
	req.CurrentUser = currentUser
	req.Per = per
	req.Page = page

	slog.Debug("hander.rows viewerReq", slog.Any("viewReq", viewReq))
	err = validateQueryParameter(where, "where")
	if err != nil {
		slog.Error("invalid character in parameter where", slog.Any("req", req), slog.Any("viewReq", viewReq), slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	err = validateQueryParameter(orderby, "orderby")
	if err != nil {
		slog.Error("invalid character in parameter orderby", slog.Any("req", req), slog.Any("viewReq", viewReq), slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	err = validateQueryParameter(search, "search")
	if err != nil {
		slog.Error("invalid character in parameter search", slog.Any("req", req), slog.Any("viewReq", viewReq), slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	rows, err := h.c.Rows(ctx.Request.Context(), req, viewReq)
	if err != nil {
		slog.Error("Failed to get dataset rows", slog.Any("req", req), slog.Any("viewReq", viewReq), slog.Any("error", err))
		httpbase.ServerError(ctx, err)
		return
	}

	httpbase.OK(ctx, rows)
}

func validateQueryParameter(parameterValue string, parameterName string) error {
	SQLInvalidSymbols := []string{";", "--", `/\*`, `\*/`}
	SQLInvalidSymbolsPattern := regexp.MustCompile(fmt.Sprintf("(?:%s)", strings.Join(SQLInvalidSymbols, "|")))

	if SQLInvalidSymbolsPattern.MatchString(parameterValue) {
		return fmt.Errorf("invalid character in %s", parameterName)
	}
	return nil
}
