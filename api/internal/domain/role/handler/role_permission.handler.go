package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
)

type RolePermissionHandler struct {
	repo repository.RolePermissionRepositoryInterface
}

func NewRolePermissionHandler(
	repo repository.RolePermissionRepositoryInterface,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		repo: repo,
	}
}

// GetRolePermissions godoc
// @Summary Get Role Permissions
// @Description Get Role Permissions
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} tools.Response{data=usecase.GetRolePermissionsOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role/{id}/permission [get]
// @Security     JWT
func (h *RolePermissionHandler) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	output, err := service.NewRolePermissionService(h.repo).GetRolePermissions(id)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Role permissions found")
	// TODO: Response
}

// CreateRolePermission godoc
// @Summary Create Role Permission
// @Description Create Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role/{id}/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = service.NewRolePermissionService(h.repo).CreateRolePermission(r.Body)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Role created")
	// TODO: Response
}

// UpdateRolePermission godoc
// @Summary Update Role Permission
// @Description Update Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "RolePermission ID"
// @Param user body usecase.UpdateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role/{id}/permission [put]
// @Security     JWT
func (h *RolePermissionHandler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = service.NewRolePermissionService(h.repo).UpdateRolePermission(id, r.Body)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Role updated")
	// TODO: Response
}
