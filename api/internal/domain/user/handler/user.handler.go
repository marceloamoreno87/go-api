package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type UserHandler struct {
	response.Responses
	service service.UserServiceInterface
}

func NewUserHandler(
	service service.UserServiceInterface,
) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user [post]
// @Security     JWT
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	err := h.service.CreateUser(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("User created")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))

}

// GetUser godoc
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=usecase.GetUserOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [get]
// @Security     JWT
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	output, err := h.service.GetUser(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("User found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))

}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetUsersOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /user [get]
// @Security     JWT
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := h.service.GetUsers(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Users found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))

}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body usecase.UpdateUserInputDTO true "User"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /user/{id} [put]
// @Security     JWT
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.UpdateUser(id, r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("User updated")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteUser(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("User deleted")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}
