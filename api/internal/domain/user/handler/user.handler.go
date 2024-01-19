package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/internal/domain/user/usecase"
	"github.com/marceloamoreno/izimoney/tools"
)

type UserHandler struct {
	HandlerTools   tools.HandlerToolsInterface
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(userRepository repository.UserRepositoryInterface, handlerTools tools.HandlerToolsInterface) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		HandlerTools:   handlerTools,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	
	var credentials usecase.GetJWTInputDTO
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get user from database
	uc := usecase.NewGetJWTUseCase(h.UserRepository)
	u, err := uc.Execute(credentials)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u.Token)

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.NewGetUserUseCase(h.UserRepository)
	u, err := uc.Execute(usecase.GetUserInputDTO{
		ID: id,
	})
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	params := usecase.GetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetUsersUseCase(h.UserRepository)
	u, err := uc.Execute(params)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user usecase.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.NewCreateUserUseCase(h.UserRepository)
	u, err := uc.Execute(user)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)

	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var user usecase.UpdateUserInputDTO
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user.ID = id
	uc := usecase.NewUpdateUserUseCase(h.UserRepository)
	u, err := uc.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.NewDeleteUserUseCase(h.UserRepository)
	err = uc.Execute(usecase.DeleteUserInputDTO{
		ID: id,
	})
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, nil)
}
