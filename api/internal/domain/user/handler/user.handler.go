package handler

import (
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type UserHandler struct {
	response.Responses
	service serviceInterface.UserServiceInterface
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: service.NewUserService(),
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user [post]
// @Security     JWT
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.CreateUser(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetUser godoc
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=usecase.GetUserOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [get]
// @Security     JWT
func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.GetUserById(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetUsersOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /user [get]
// @Security     JWT
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset := helper.GetLimitAndOffset(r)
	output, err := h.service.GetUsers(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body usecase.UpdateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [put]
// @Security     JWT
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.UpdateUser(helper.GetID(r), r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.DeleteUser(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
