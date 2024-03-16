package handler

import (
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type PermissionHandler struct {
	response.Responses
	service serviceInterface.PermissionServiceInterface
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{
		service: service.NewPermissionService(),
	}
}

// GetPermission godoc
// @Summary Get Permission
// @Description Get Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} response.Response{data=usecase.GetPermissionOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission/{id} [get]
// @Security     JWT
func (h *PermissionHandler) GetPermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.GetPermission(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetPermissionsOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission [get]
// @Security     JWT
func (h *PermissionHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {
	limit, offset := helper.GetLimitAndOffset(r)
	output, err := h.service.GetPermissions(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// CreateRole godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param role body usecase.CreatePermissionInputDTO true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [post]
// @Security     JWT
func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.CreatePermission(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdateRole godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param role body usecase.UpdatePermissionInputDTO true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission/{id} [put]
// @Security     JWT
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.UpdatePermission(helper.GetID(r), r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// DeletePermission godoc
// @Summary Delete Permission
// @Description Delete Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /permission/{id} [delete]
// @Security     JWT
func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.DeletePermission(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
