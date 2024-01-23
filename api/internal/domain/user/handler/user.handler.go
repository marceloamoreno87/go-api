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

// GetUser godoc
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} tools.Response{data=entity.User}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user/{id} [get]
// @Security     JWT
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

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} tools.Response{data=[]entity.User}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user [get]
// @Security     JWT
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
	h.HandlerTools.ResponseJSON(w, http.StatusOK, u.Users)
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateUserInputDTO true "User"
// @Success 200 {object} tools.Response{data=entity.User}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user [post]
// @Security     JWT
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

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body usecase.UpdateUserInputDTO true "User"
// @Success 200 {object} tools.Response{data=entity.User}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user/{id} [put]
// @Security     JWT
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

	uc := usecase.NewUpdateUserUseCase(h.UserRepository, id)
	u, err := uc.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} tools.Response{data=usecase.DeleteUserOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.NewDeleteUserUseCase(h.UserRepository)
	u, err := uc.Execute(usecase.DeleteUserInputDTO{
		ID: id,
	})
	if err != nil {
		h.HandlerTools.ResponseErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.HandlerTools.ResponseJSON(w, http.StatusOK, u)
}
