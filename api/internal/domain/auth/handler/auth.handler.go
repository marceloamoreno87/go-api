package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/auth/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type AuthHandler struct {
	repo repository.UserRepositoryInterface
}

func NewAuthHandler(
	repo repository.UserRepositoryInterface,
) *AuthHandler {
	return &AuthHandler{
		repo: repo,
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
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	token, err := service.NewAuthService(h.repo).Login(r.Body)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	slog.Info("Login success")
	// TODO: Response

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
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	token, err := service.NewAuthService(h.repo).Refresh(r.Body)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Login refreshed")
	// TODO: Response
}
