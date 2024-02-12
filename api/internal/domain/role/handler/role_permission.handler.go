package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type RolePermissionHandler struct {
	HandlerTools             api.HandlerToolsInterface
	RolePermissionRepository repository.RolePermissionRepositoryInterface
}

func NewRolePermissionHandler(
	RolePermissionRepository repository.RolePermissionRepositoryInterface,
	handlerTools api.HandlerToolsInterface,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		RolePermissionRepository: RolePermissionRepository,
		HandlerTools:             handlerTools,
	}
}

// GetRolePermissions godoc
// @Summary Get Role Permissions
// @Description Get Role Permissions
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Success 200 {object} api.Response{data=entity.RolePermission}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id}/permission [get]
// @Security     JWT
func (h *RolePermissionHandler) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetRolePermissionsUseCase(h.RolePermissionRepository)
	rolePermissions, err := uc.Execute(usecase.GetRolePermissionsInputDTO{
		RoleID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role permissions get", "role permissions", rolePermissions)
	h.HandlerTools.ResponseJSON(w, rolePermissions)

}

// CreateRolePermission godoc
// @Summary Create Role Permission
// @Description Create Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param user body usecase.RolePermissionInputDTO true "RolePermission"
// @Success 200 {object} api.Response{data=entity.RolePermission}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
}

// UpdateRolePermission godoc
// @Summary Update Role Permission
// @Description Update Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "RolePermission ID"
// @Param user body usecase.RolePermissionInputDTO true "RolePermission"
// @Success 200 {object} api.Response{data=entity.RolePermission}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/permission/{id} [put]
// @Security     JWT
func (h *RolePermissionHandler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
}
