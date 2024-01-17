package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/internal/domain/user/usecase"
	"github.com/marceloamoreno/izimoney/tools"
)

type UserHandler struct {
	HandlerTools   *tools.HandlerTools
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(userRepository repository.UserRepositoryInterface, handlerTools *tools.HandlerTools) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		HandlerTools:   handlerTools,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewGetUserUseCase(h.UserRepository)
	u, err := uc.Execute(usecase.GetUserInputDTO{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.HandlerTools.ResponseJSON(w, "Success", http.StatusOK, u)

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := usecase.GetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetUsersUseCase(h.UserRepository)
	u, err := uc.Execute(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.HandlerTools.ResponseJSON(w, "Success", http.StatusOK, u)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user usecase.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewCreateUserUseCase(h.UserRepository)
	u, err := uc.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.HandlerTools.ResponseJSON(w, "Created user successfully", http.StatusOK, u)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user usecase.UpdateUserInputDTO
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id
	uc := usecase.NewUpdateUserUseCase(h.UserRepository)
	u, err := uc.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.HandlerTools.ResponseJSON(w, "Updated user successfully", http.StatusOK, u)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewDeleteUserUseCase(h.UserRepository)
	err = uc.Execute(usecase.DeleteUserInputDTO{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.HandlerTools.ResponseJSON(w, "Deleted user successfully", http.StatusOK, nil)
}
