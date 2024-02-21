package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type RoleHandler struct {
	HandlerTools api.HandlerToolsInterface
	repo         repository.RoleRepositoryInterface
}

func NewRoleHandler(
	repo repository.RoleRepositoryInterface,
	handlerTools api.HandlerToolsInterface,
) *RoleHandler {
	return &RoleHandler{
		repo:         repo,
		HandlerTools: handlerTools,
	}
}

// GetRole godoc
// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} api.Response{data=usecase.GetRoleOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id} [get]
// @Security     JWT
func (h *RoleHandler) GetRole(w http.ResponseWriter, r *http.Request) {

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetRoleUseCase(h.repo)
	role, err := uc.Execute(usecase.GetRoleInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role get", "roles", role)
	h.HandlerTools.ResponseJSON(w, role)

}

// GetRoles godoc
// @Summary Get Roles
// @Description Get Roles
// @Tags Role
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} api.Response{data=[]usecase.GetRolesOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role [get]
// @Security     JWT
func (h *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	dto := usecase.GetRolesInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetRolesUseCase(h.repo)
	role, err := uc.Execute(dto)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Roles getting", "roles", role)
	h.HandlerTools.ResponseJSON(w, role)
}

// CreateRole godoc
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateRoleInputDTO true "Role"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role [post]
// @Security     JWT
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {

	var input usecase.CreateRoleInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewCreateRoleUseCase(h.repo)
	err = uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role created")
	h.HandlerTools.ResponseJSON(w, nil)

}

// UpdateRole godoc
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role body usecase.UpdateUserInputDTO true "Role"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/{id} [put]
// @Security     JWT
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	var input usecase.UpdateRoleInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewUpdateRoleUseCase(h.repo, id)
	err = uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role updated")
	h.HandlerTools.ResponseJSON(w, nil)
}

// DeleteRole godoc
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} api.Response{data=nil}
// @Failure 400 {object} api.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /role/{id} [delete]
// @Security     JWT
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewDeleteRoleUseCase(h.repo)
	err = uc.Execute(usecase.DeleteRoleInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role deleted")
	h.HandlerTools.ResponseJSON(w, nil)
}
