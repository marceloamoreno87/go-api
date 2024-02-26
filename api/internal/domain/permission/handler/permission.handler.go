package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/service"
)

type PermissionHandler struct {
	repo repository.PermissionRepositoryInterface
}

func NewPermissionHandler(
	repo repository.PermissionRepositoryInterface,
) *PermissionHandler {
	return &PermissionHandler{
		repo: repo,
	}
}

// GetPermission godoc
// @Summary Get Permission
// @Description Get Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} tools.Response{data=usecase.GetPermissionOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /permission/{id} [get]
// @Security     JWT
func (h *PermissionHandler) GetPermission(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	output, err := service.NewPermissionService(h.repo).GetPermission(id)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Permission found")
	// TODO: Response

}

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} tools.Response{data=[]usecase.GetPermissionsOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /permission [get]
// @Security     JWT
func (h *PermissionHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := service.NewPermissionService(h.repo).GetPermissions(limit, offset)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Users found")
	// TODO: Response

}

// CreateRole godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateRoleInputDTO true "Permission"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role [post]
// @Security     JWT
func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewPermissionService(h.repo).CreatePermission(r.Body)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
			return
		}
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

	slog.Info("User created")
	// TODO: Response

}

// UpdateRole godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param role body usecase.UpdateUserInputDTO true "Permission"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /permission/{id} [put]
// @Security     JWT
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewPermissionService(h.repo).UpdatePermission(id, r.Body)

	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
			return
		}
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

	slog.Info("User updated")
	// TODO: Response
}

// DeletePermission godoc
// @Summary Delete Permission
// @Description Delete Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /permission/{id} [delete]
// @Security     JWT
func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewPermissionService(h.repo).DeletePermission(id)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
		}
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

	slog.Info("User deleted")
	// TODO: Response

}
