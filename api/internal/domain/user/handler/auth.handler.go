package handler

import (
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AuthHandler struct {
	response.Responses
	service serviceInterface.AuthServiceInterface
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: service.NewAuthService(),
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param auth body service.RequestLoginInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.CreateAuthOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.Login(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetRefreshJWT godoc
// @Summary Get Refresh JWT
// @Description Get Refresh JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param auth body service.RequestRefreshTokenInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.CreateAuthOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.RefreshToken(r.Body)
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
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param auth body service.RequestForgotPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	err := h.service.ForgotPassword(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// UpdateUserPassword godoc
// @Summary Update User Password
// @Description Update User Password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param auth body service.RequestUpdateUserPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/update-password [post]
func (h *AuthHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	err := h.service.UpdateUserPassword(r.Body)
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
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param auth body service.RequestVerifyUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/verify-user [post]
func (h *AuthHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	err := h.service.VerifyUser(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}
