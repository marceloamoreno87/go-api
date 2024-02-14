package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type UserHandler struct {
	HandlerTools   api.HandlerToolsInterface
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(userRepository repository.UserRepositoryInterface, handlerTools api.HandlerToolsInterface) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		HandlerTools:   handlerTools,
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateUserInputDTO true "User"
// @Success 200 {object} api.Response{data=entity.User}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /user [post]
// @Security     JWT
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var input usecase.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewCreateUserUseCase(h.UserRepository)
	err = uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("User created", "user")
	h.HandlerTools.ResponseJSON(w, nil)

}

// GetUser godoc
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} api.Response{data=entity.User}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /user/{id} [get]
// @Security     JWT
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewGetUserUseCase(h.UserRepository)
	u, err := uc.Execute(usecase.GetUserInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("User get", "users", u)
	h.HandlerTools.ResponseJSON(w, u)

}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} api.Response{data=[]entity.User}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /user [get]
// @Security     JWT
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.HandlerTools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	input := usecase.GetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetUsersUseCase(h.UserRepository)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Users getting", "users", u)
	h.HandlerTools.ResponseJSON(w, u)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body usecase.UpdateUserInputDTO true "User"
// @Success 200 {object} api.Response{data=entity.User}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /user/{id} [put]
// @Security     JWT
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	var input usecase.UpdateUserInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewUpdateUserUseCase(h.UserRepository, id)
	err = uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("User updated")
	h.HandlerTools.ResponseJSON(w, nil)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} api.Response{data=usecase.DeleteUserOutputDTO}
// @Failure 400 {object} api.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.HandlerTools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewDeleteUserUseCase(h.UserRepository)
	err = uc.Execute(usecase.DeleteUserInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("User deleted")
	h.HandlerTools.ResponseJSON(w, nil)
}
