package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	UserVerify(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	response.Responses
	service service.UserServiceInterface
}

func NewAuthHandler(service service.UserServiceInterface) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body usecase.LoginInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.LoginOutputDTO}
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
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body usecase.RegisterInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.RegisterOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.Register(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetVerify godoc
// @Summary Get Verify
// @Description Get Verify
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body usecase.UserVerifyInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.RegisterOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/verify [post]
func (h *AuthHandler) UserVerify(w http.ResponseWriter, r *http.Request) {
	err := h.service.UserVerify(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// GetForgotPassword godoc
// @Summary Get Forgot Password
// @Description Get Forgot Password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body usecase.ForgotPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.ForgotPasswordOutputDTO}
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

// GetUpdatePasswordUser godoc
// @Summary Get Update Password User
// @Description Get Update Password User
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body usecase.UpdatePasswordUserInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.UpdatePasswordUserOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/update-password [patch]
func (h *AuthHandler) UpdatePasswordUser(w http.ResponseWriter, r *http.Request) {
	err := h.service.UpdatePasswordUser(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}
