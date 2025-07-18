package handler

import (
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/common/config"
	"opencsg.com/csghub-server/common/errorx"
	"opencsg.com/csghub-server/common/types"
	"opencsg.com/csghub-server/component"
)

type SyncClientSettingHandler struct {
	c component.SyncClientSettingComponent
}

func NewSyncClientSettingHandler(config *config.Config) (*SyncClientSettingHandler, error) {
	c, err := component.NewSyncClientSettingComponent(config)
	if err != nil {
		return nil, err
	}
	return &SyncClientSettingHandler{
		c: c,
	}, nil
}

// CreateSyncClientSetting  godoc
// @Security     ApiKey
// @Summary      Create sync client setting or update an existing sync client setting
// @Description  Create sync client setting or update an existing sync client setting
// @Tags         Sync
// @Accept       json
// @Produce      json
// @Param        body   body  types.CreateSyncClientSettingReq true "body"
// @Success      200  {object}  types.Response{data=database.SyncClientSetting} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /sync/client_setting [post]
func (h *SyncClientSettingHandler) Create(ctx *gin.Context) {
	var req types.CreateSyncClientSettingReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Bad request format", "error", err)
		httpbase.BadRequest(ctx, err.Error())
		return
	}
	req.CurrentUser = httpbase.GetCurrentUser(ctx)
	if req.CurrentUser == "" {
		httpbase.UnauthorizedError(ctx, errorx.ErrUserNotFound)
		return
	}
	ms, err := h.c.Create(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, errorx.ErrUnauthorized) {
			httpbase.UnauthorizedError(ctx, err)
			return
		}
		slog.Error("Failed to create sync client setting", "error", err)
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, ms)
}

// GetSyncClientSetting  godoc
// @Security     ApiKey
// @Summary      Get sync client setting
// @Description  Get sync client setting
// @Tags         Sync
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=database.SyncClientSetting} "OK"
// @Failure      400  {object}  types.APIBadRequest "Bad request"
// @Failure      500  {object}  types.APIInternalServerError "Internal server error"
// @Router       /sync/client_setting [get]
func (h *SyncClientSettingHandler) Show(ctx *gin.Context) {
	currentUser := httpbase.GetCurrentUser(ctx)
	if currentUser == "" {
		httpbase.UnauthorizedError(ctx, errorx.ErrUserNotFound)
		return
	}
	ms, err := h.c.Show(ctx.Request.Context(), currentUser)
	if err != nil {
		if errors.Is(err, errorx.ErrUnauthorized) {
			httpbase.UnauthorizedError(ctx, err)
			return
		}
		slog.Error("Failed to find sync client setting", "error", err)
		httpbase.ServerError(ctx, err)
		return
	}
	httpbase.OK(ctx, ms)
}
