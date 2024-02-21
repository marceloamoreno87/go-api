package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type PermissionHandler struct {
	handlerTools api.HandlerToolsInterface
	repo         repository.PermissionRepositoryInterface
}

func NewPermissionHandler(
	repo repository.PermissionRepositoryInterface,
	handlerTools api.HandlerToolsInterface,
) *PermissionHandler {
	return &PermissionHandler{
		repo:         repo,
		handlerTools: handlerTools,
	}
}

// GetPermission godoc
// @Summary Get Permission
// @Description Get Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} api.Response{data=usecase.GetPermissionOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /permission/{id} [get]
// @Security     JWT
func (h *PermissionHandler) GetPermission(w http.ResponseWriter, r *http.Request) {

	id, err := h.handlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetPermissionUseCase(h.repo)
	permission, err := uc.Execute(usecase.GetPermissionInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission get", "permissions", permission)
	h.handlerTools.ResponseJSON(w, permission)

}

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} api.Response{data=[]usecase.GetPermissionsOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /permission [get]
// @Security     JWT
func (h *PermissionHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.handlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
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
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permissions getting", "permissions", permission)
	h.handlerTools.ResponseJSON(w, permission)
}

// CreateRole godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateRoleInputDTO true "Permission"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role [post]
// @Security     JWT
func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {

	var dto usecase.CreatePermissionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewCreatePermissionUseCase(h.repo)
	err = uc.Execute(dto)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err2.Error()))
		}
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission created")
	h.handlerTools.ResponseJSON(w, nil)

}

// UpdateRole godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param role body usecase.UpdateUserInputDTO true "Permission"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /permission/{id} [put]
// @Security     JWT
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	id, err := h.handlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	var dto usecase.UpdatePermissionInputDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	uc := usecase.NewUpdatePermissionUseCase(h.repo, id)
	err = uc.Execute(dto)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err2.Error()))
		}
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission updated")
	h.handlerTools.ResponseJSON(w, nil)
}

// DeletePermission godoc
// @Summary Delete Permission
// @Description Delete Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /permission/{id} [delete]
// @Security     JWT
func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	id, err := h.handlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
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
			h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err2.Error()))
		}
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission deleted")
	h.handlerTools.ResponseJSON(w, nil)
}
