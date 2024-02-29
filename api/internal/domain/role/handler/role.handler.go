package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type RoleHandler struct {
	response.Responses
	service service.RoleServiceInterface
}

func NewRoleHandler(
	service service.RoleServiceInterface,
) *RoleHandler {
	return &RoleHandler{
		service: service,
	}
}

// GetRole godoc
// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response{data=usecase.GetRoleOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id} [get]
// @Security     JWT
func (h *RoleHandler) GetRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	output, err := h.service.GetRole(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))
}

// GetRoles godoc
// @Summary Get Roles
// @Description Get Roles
// @Tags Role
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetRolesOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [get]
// @Security     JWT
func (h *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := h.service.GetRoles(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Roles found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))

}

// CreateRole godoc
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateRoleInputDTO true "Role"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [post]
// @Security     JWT
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {

	err := h.service.CreateRole(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role created")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}

// UpdateRole godoc
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role body usecase.UpdateRoleInputDTO true "Role"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id} [put]
// @Security     JWT
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.service.UpdateRole(id, r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role updated")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))

}

// DeleteRole godoc
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /role/{id} [delete]
// @Security     JWT
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.service.DeleteRole(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Role deleted")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}
