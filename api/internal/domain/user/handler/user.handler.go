package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/domain/user/usecase"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
	"github.com/marceloamoreno/izimoney/tools"
)

type UserHandler struct {
	HandlerTools   *tools.HandlerTools
	UserRepository *db.Queries
}

func NewUserHandler(userRepository *db.Queries, handlerTools *tools.HandlerTools) *UserHandler {
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
	user, err := uc.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	GetUsersParams := db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetUsersUseCase(h.UserRepository)
	users, err := uc.Execute(GetUsersParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var CreateUserParams db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&CreateUserParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewCreateUserUseCase(h.UserRepository)
	user, err := uc.Execute(CreateUserParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UpdateUserParams := db.UpdateUserParams{
		ID: id,
	}

	err = json.NewDecoder(r.Body).Decode(&UpdateUserParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewUpdateUserUseCase(h.UserRepository)
	user, err := uc.Execute(UpdateUserParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc := usecase.NewDeleteUserUseCase(h.UserRepository)
	err = uc.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
