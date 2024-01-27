package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/tools"
)

type AuthHandler struct {
	HandlerTools   tools.HandlerToolsInterface
	UserRepository repository.UserRepositoryInterface
}

func NewAuthHandler(userRepository repository.UserRepositoryInterface, handlerTools tools.HandlerToolsInterface) *AuthHandler {
	return &AuthHandler{
		UserRepository: userRepository,
		HandlerTools:   handlerTools,
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

	var credentials usecase.GetJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	uc := usecase.NewGetJWTUseCase(h.UserRepository)
	u, err := uc.Execute(credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	slog.Info("Token generated", "token", u)
	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)

}
