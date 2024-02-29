package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type RolePermissionHandler struct {
	response.Responses
	service service.RolePermissionServiceInterface
}

func NewRolePermissionHandler(
	service service.RolePermissionServiceInterface,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		service: service,
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
	id := chi.URLParam(r, "id")

	output, err := h.service.GetRolePermissions(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role permissions found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))
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

	err := h.service.CreateRolePermission(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role permission created")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))

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
func (h *RolePermissionHandler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.UpdateRolePermission(id, r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role permission updated")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}
