package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
	"github.com/marceloamoreno/goapi/internal/shared/validate"
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
// @Param user body request.RequestCreateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user [post]
// @Security     JWT
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestCreateUserInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.CreateUser(input)
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
// @Success 200 {object} response.Response{data=request.RequestGetUserInputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [get]
// @Security     JWT
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestGetUserInputDTO{
		ID: helper.GetID(r),
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.GetUser(input)
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
	input := request.RequestGetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.GetUsers(input)
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
// @Param user body request.RequestUpdateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [put]
// @Security     JWT
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestUpdateUserInputDTO{
		ID: helper.GetID(r),
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := validate.NewValidator(input).Validate()
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	output, err := h.service.UpdateUser(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdateUserPassword godoc
// @Summary Update User Password
// @Description Update User Password
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.RequestUpdateUserPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/update-password [post]
func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	input := request.RequestUpdateUserPasswordInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := h.service.UpdateUserPassword(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
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
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestDeleteUserInputDTO{
		ID: helper.GetID(r),
	}
	output, err := h.service.DeleteUser(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// ForgotPassword godoc
// @Summary Forgot Password
// @Description Forgot Password
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.RequestForgotPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/forgot-password [post]
func (h *UserHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	input := request.RequestForgotPasswordInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := h.service.ForgotPassword(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// VerifyUser godoc
// @Summary Verify User
// @Description Verify User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.RequestVerifyUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/verify-user [post]
func (h *UserHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestVerifyUserInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	err := h.service.VerifyUser(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// RegisterUser godoc
// @Summary Register User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.RequestCreateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/register [post]
// @Security     JWT
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	input := request.RequestCreateUserInputDTO{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.CreateUser(input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
