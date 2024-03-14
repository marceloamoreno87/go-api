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
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	response.Responses
	service service.AuthServiceInterface
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
