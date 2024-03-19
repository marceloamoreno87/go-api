package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type RolePermissionHandler struct {
	response.Responses
	service serviceInterface.RolePermissionServiceInterface
}

func NewRolePermissionHandler() *RolePermissionHandler {
	return &RolePermissionHandler{
		service: service.NewRolePermissionService(),
	}
}

// GetRolePermissions godoc
// @Summary Get Role Permissions
// @Description Get Role Permissions
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response{data=usecase.GetRolePermissionsOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [get]
// @Security     JWT
func (h *RolePermissionHandler) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	input := service.RequestGetRolePermissionInputDTO{
		RoleID: helper.GetID(r),
	}
	output, err := h.service.GetRolePermissions(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	slog.Info("Role permissions found")
	h.SendResponse(w, h.NewResponse(output))
}

// CreateRolePermission godoc
// @Summary Create Role Permission
// @Description Create Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param role_permission body service.RequestCreateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	input := service.RequestCreateRolePermissionInputDTO{
		RoleID: helper.GetID(r),
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.CreateRolePermission(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	slog.Info("Role permission created")
	h.SendResponse(w, h.NewResponse(output))

}

// UpdateRolePermission godoc
// @Summary Update Role Permission
// @Description Update Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role_permission body service.RequestUpdateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [put]
// @Security     JWT
func (h *RolePermissionHandler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
	input := service.RequestUpdateRolePermissionInputDTO{
		RoleID: helper.GetID(r),
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.UpdateRolePermission(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// DeleteRolePermissionByRoleID godoc
// @Summary Delete Role Permission By Role ID
// @Description Delete Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [delete]
// @Security     JWT
func (h *RolePermissionHandler) DeleteRolePermissionByRoleID(w http.ResponseWriter, r *http.Request) {
	input := service.RequestDeleteRolePermissionByRoleIDInputDTO{
		RoleID: helper.GetID(r),
	}
	output, err := h.service.DeleteRolePermissionByRoleID(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
