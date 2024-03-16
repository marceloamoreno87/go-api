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
	output, err := h.service.GetRolePermissions(helper.GetID(r))
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
// @Param user body usecase.CreateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.CreateRolePermission(r.Body)
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
// @Param id path string true "RolePermission ID"
// @Param user body usecase.UpdateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id}/permission [put]
// @Security     JWT
func (h *RolePermissionHandler) DeleteRolePermission(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.DeleteRolePermission(helper.GetID(r), r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
