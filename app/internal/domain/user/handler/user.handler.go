package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marceloamoreno/izimoney/internal/db"
	"github.com/marceloamoreno/izimoney/internal/domain/user/usecase"
)

type UserHandler struct {
	UserRepository *db.Queries
}

func NewUserHandler(userRepository *db.Queries) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
}
