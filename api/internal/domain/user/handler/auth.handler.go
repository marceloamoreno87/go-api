package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/config"
	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	response.Responses
	service serviceInterface.AuthServiceInterface
}

func NewAuthHandler() *AuthHandler {
	DB := config.NewSqlc(config.NewDatabase())
	return &AuthHandler{
		service: service.NewAuthService(DB),
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
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.RefreshToken(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
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

// UpdateUserPassword godoc
// @Summary Update User Password
// @Description Update User Password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body usecase.UpdateUserPasswordInputDTO true "User"
// @Success 200 {object} response.Response{data=usecase.UpdateUserPasswordOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/{id}/update-password [patch]
func (h *AuthHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.UpdateUserPassword(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
