package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
	"github.com/marceloamoreno/goapi/internal/shared/validate"
)

type PermissionHandler struct {
	response.Responses
	service service.PermissionService
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{
		service: *service.NewPermissionService(),
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
	input := request.RequestGetPermission{
		ID: helper.GetID(r),
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.GetPermission(r.Context(), input)
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
	input := request.RequestGetPermissions{
		Limit:  limit,
		Offset: offset,
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.GetPermissions(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// CreatePermission godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param permission body request.RequestCreatePermission true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission [post]
// @Security     JWT
func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	input := request.RequestCreatePermission{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.CreatePermission(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdatePermission godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param permission body request.RequestUpdatePermission true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission/{id} [put]
// @Security     JWT
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	input := request.RequestUpdatePermission{
		ID: helper.GetID(r),
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.UpdatePermission(r.Context(), input)
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
	input := request.RequestDeletePermission{
		ID: helper.GetID(r),
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.DeletePermission(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
