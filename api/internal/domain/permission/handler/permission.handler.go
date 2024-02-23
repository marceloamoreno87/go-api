package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/pkg/tools"
)

type PermissionHandler struct {
	tools tools.HandlerToolsInterface
	repo  repository.PermissionRepositoryInterface
}

func NewPermissionHandler(
	repo repository.PermissionRepositoryInterface,
) *PermissionHandler {
	return &PermissionHandler{
		repo:  repo,
		tools: tools.NewHandlerTools(),
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

	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	uc := usecase.NewGetPermissionUseCase(h.repo)
	permission, err := uc.Execute(usecase.GetPermissionInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	slog.Info("Permission get", "permissions", permission)
	h.tools.ResponseJSON(w, permission)

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
	limit, offset, err := h.tools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	dto := usecase.GetPermissionsInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetPermissionsUseCase(h.repo)
	permission, err := uc.Execute(dto)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	slog.Info("Permissions getting", "permissions", permission)
	h.tools.ResponseJSON(w, permission)
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

	var dto usecase.CreatePermissionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	uc := usecase.NewCreatePermissionUseCase(h.repo)
	err = uc.Execute(dto)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.MountError(err2, http.StatusBadRequest, "BAR_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	slog.Info("Permission created")
	h.tools.ResponseJSON(w, nil)

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
	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	var dto usecase.UpdatePermissionInputDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	uc := usecase.NewUpdatePermissionUseCase(h.repo, id)
	err = uc.Execute(dto)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.MountError(err2, http.StatusBadRequest, "BAR_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	slog.Info("Permission updated")
	h.tools.ResponseJSON(w, nil)
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
	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	uc := usecase.NewDeletePermissionUseCase(h.repo)
	err = uc.Execute(usecase.DeletePermissionInputDTO{
		ID: id,
	})
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.MountError(err2, http.StatusBadRequest, "BAR_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	slog.Info("Permission deleted")
	h.tools.ResponseJSON(w, nil)
}
