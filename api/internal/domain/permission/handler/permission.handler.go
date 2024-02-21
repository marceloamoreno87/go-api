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
	HandlerTools api.HandlerToolsInterface
	Repo         repository.PermissionRepositoryInterface
}

func NewPermissionHandler(Repo repository.PermissionRepositoryInterface, handlerTools api.HandlerToolsInterface) *PermissionHandler {
	return &PermissionHandler{
		Repo:         Repo,
		HandlerTools: handlerTools,
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

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetPermissionUseCase(h.Repo)
	permission, err := uc.Execute(usecase.GetPermissionInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission get", "permissions", permission)
	h.HandlerTools.ResponseJSON(w, permission)

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
	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	dto := usecase.GetPermissionsInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetPermissionsUseCase(h.Repo)
	permission, err := uc.Execute(dto)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permissions getting", "permissions", permission)
	h.HandlerTools.ResponseJSON(w, permission)
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
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewCreatePermissionUseCase(h.Repo)
	err = uc.Execute(dto)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission created")
	h.HandlerTools.ResponseJSON(w, nil)

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
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	var dto usecase.UpdatePermissionInputDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewUpdatePermissionUseCase(h.Repo, id)
	err = uc.Execute(dto)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission updated")
	h.HandlerTools.ResponseJSON(w, nil)
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
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewDeletePermissionUseCase(h.Repo)
	err = uc.Execute(usecase.DeletePermissionInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Permission deleted")
	h.HandlerTools.ResponseJSON(w, nil)
}
