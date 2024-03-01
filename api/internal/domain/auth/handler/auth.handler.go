package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/auth/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AuthHandler struct {
	response.Responses
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
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
// @Param token body usecase.GetJWTInputDTO true "Token"
// @Success 200 {object} response.Response{data=usecase.GetJWTOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	token, err := h.service.Login(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	slog.Info("Login success")
	h.SendResponse(w, h.NewResponse(token))
}

// GetRefreshJWT godoc
// @Summary Get Refresh JWT
// @Description Get Refresh JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param token body usecase.GetRefreshJWTInputDTO true "Token"
// @Success 200 {object} response.Response{data=usecase.GetRefreshJWTOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	token, err := h.service.Refresh(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}

	slog.Info("Login refreshed")
	h.SendResponse(w, h.NewResponse(token))
}
