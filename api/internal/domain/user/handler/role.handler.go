package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type RoleHandler struct {
	response.Responses
	service serviceInterface.RoleServiceInterface
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{
		service: service.NewRoleService(),
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
	input := service.RequestGetRoleInputDTO{
		ID: helper.GetID(r),
	}
	output, err := h.service.GetRole(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
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
	limit, offset := helper.GetLimitAndOffset(r)
	input := service.RequestGetRolesInputDTO{
		Limit:  limit,
		Offset: offset,
	}
	output, err := h.service.GetRoles(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// CreateRole godoc
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param role body service.RequestCreateRoleInputDTO true "Role"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [post]
// @Security     JWT
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	input := service.RequestCreateRoleInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.CreateRole(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdateRole godoc
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role body service.RequestUpdateRoleInputDTO true "Role"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id} [put]
// @Security     JWT
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	input := service.RequestUpdateRoleInputDTO{
		ID: helper.GetID(r),
	}
	output, err := h.service.UpdateRole(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// DeleteRole godoc
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role/{id} [delete]
// @Security     JWT
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	input := service.RequestDeleteRoleInputDTO{
		ID: helper.GetID(r),
	}
	output, err := h.service.DeleteRole(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
