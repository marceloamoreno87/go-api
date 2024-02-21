package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type RolePermissionHandler struct {
	handlerTools api.HandlerToolsInterface
	repo         repository.RolePermissionRepositoryInterface
}

func NewRolePermissionHandler(
	repo repository.RolePermissionRepositoryInterface,
	handlerTools api.HandlerToolsInterface,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		repo:         repo,
		handlerTools: handlerTools,
	}
}

// GetRolePermissions godoc
// @Summary Get Role Permissions
// @Description Get Role Permissions
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} api.Response{data=usecase.GetRolePermissionsOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id}/permission [get]
// @Security     JWT
func (h *RolePermissionHandler) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	id, err := h.handlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetRolePermissionsUseCase(h.repo)
	rolePermissions, err := uc.Execute(usecase.GetRolePermissionsInputDTO{
		RoleID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.handlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role permissions get", "role permissions", rolePermissions)
	h.handlerTools.ResponseJSON(w, rolePermissions)
}

// CreateRolePermission godoc
// @Summary Create Role Permission
// @Description Create Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id}/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateRolePermissionInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
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
	uc := usecase.NewCreateRolePermissionUseCase(h.repo)
	err = uc.Execute(input)
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
	h.handlerTools.ResponseJSON(w, nil)
}

// UpdateRolePermission godoc
// @Summary Update Role Permission
// @Description Update Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param id path string true "RolePermission ID"
// @Param user body usecase.UpdateRolePermissionInputDTO true "RolePermission"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id}/permission [put]
// @Security     JWT
func (h *RolePermissionHandler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateRolePermissionInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
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
	uc := usecase.NewUpdateRolePermissionUseCase(h.repo)
	err = uc.Execute(input)
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
	h.handlerTools.ResponseJSON(w, nil)
}
