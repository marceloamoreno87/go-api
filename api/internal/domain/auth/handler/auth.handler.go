package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/domain/auth/usecase"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/tools"
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
// @Param credentials body GetJWTInputDTO true "Credentials"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /auth/jwt [post]

func (h *AuthHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	var credentials usecase.GetJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.NewGetJWTUseCase(h.UserRepository)
	u, err := uc.Execute(credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u.Token)

}
