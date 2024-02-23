package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/pkg/tools"
)

type AuthHandler struct {
	tools tools.HandlerToolsInterface
	repo  repository.UserRepositoryInterface
}

func NewAuthHandler(
	repo repository.UserRepositoryInterface,
) *AuthHandler {
	return &AuthHandler{
		repo:  repo,
		tools: tools.NewHandlerTools(),
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param credentials body usecase.GetJWTInputDTO true "Credentials"
// @Success 200 {object} tools.Response{data=usecase.GetJWTOutputDTO}
// @Failure 400 {object} tools.ResponseError
// @Router /auth/token [post]
func (h *AuthHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	var input usecase.GetJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusBadRequest, "BAR_REQUEST"))
		return
	}
	uc := usecase.NewGetJWTUseCase(h.repo)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusUnauthorized, "NOT_AUTHORIZED"))
		return
	}
	slog.Info("Token generated", "token", u)
	h.tools.ResponseJSON(w, u)

}

// GetRefreshJWT godoc
// @Summary Get Refresh JWT
// @Description Get Refresh JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param token body usecase.GetRefreshJWTInputDTO true "Token"
// @Success 200 {object} tools.Response{data=usecase.GetRefreshJWTOutputDTO}
// @Failure 400 {object} tools.ResponseError
// @Router /auth/token/refresh [post]
func (h *AuthHandler) GetRefreshJWT(w http.ResponseWriter, r *http.Request) {

	var input usecase.GetRefreshJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusUnauthorized, "NOT_AUTHORIZED"))
		return
	}

	uc := usecase.NewGetRefreshJWTUseCase(h.repo)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.MountError(err, http.StatusUnauthorized, "NOT_AUTHORIZED"))
		return
	}
	slog.Info("Token refreshed", "token", u)
	h.tools.ResponseJSON(w, u)

}
