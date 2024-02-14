package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type AuthHandler struct {
	HandlerTools   api.HandlerToolsInterface
	UserRepository repository.UserRepositoryInterface
}

func NewAuthHandler(UserRepository repository.UserRepositoryInterface, handlerTools api.HandlerToolsInterface) *AuthHandler {
	return &AuthHandler{
		HandlerTools:   handlerTools,
		UserRepository: UserRepository,
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param credentials body usecase.GetJWTInputDTO true "Credentials"
// @Success 200 {object} api.Response{data=usecase.GetJWTOutputDTO}
// @Failure 400 {object} api.ResponseError
// @Router /auth/token [post]
func (h *AuthHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	var input usecase.GetJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetJWTUseCase(h.UserRepository)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NOT_AUTHORIZED)
		return
	}
	slog.Info("Token generated", "token", u)
	h.HandlerTools.ResponseJSON(w, u)

}

// GetRefreshJWT godoc
// @Summary Get Refresh JWT
// @Description Get Refresh JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param token body usecase.GetRefreshJWTInputDTO true "Token"
// @Success 200 {object} api.Response{data=usecase.GetRefreshJWTOutputDTO}
// @Failure 400 {object} api.ResponseError
// @Router /auth/token/refresh [post]
func (h *AuthHandler) GetRefreshJWT(w http.ResponseWriter, r *http.Request) {

	var input usecase.GetRefreshJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetRefreshJWTUseCase(h.UserRepository)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NOT_AUTHORIZED)
		return
	}
	slog.Info("Token refreshed", "token", u)
	h.HandlerTools.ResponseJSON(w, u)

}
