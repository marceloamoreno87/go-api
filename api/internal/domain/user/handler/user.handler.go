package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/pkg/tools"
)

type UserHandler struct {
	repo  repository.UserRepositoryInterface
	tools tools.HandlerToolsInterface
}

func NewUserHandler(
	repo repository.UserRepositoryInterface,
) *UserHandler {
	return &UserHandler{
		repo:  repo,
		tools: tools.NewHandlerTools(),
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body usecase.CreateUserInputDTO true "User"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user [post]
// @Security     JWT
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var input usecase.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	uc := usecase.NewCreateUserUseCase(h.repo)
	err = uc.Execute(input)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err2.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	slog.Info("User created")
	h.tools.ResponseJSON(w, h.tools.NewResponse(nil, http.StatusOK))
}

// GetUser godoc
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} tools.Response{data=usecase.GetUserOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user/{id} [get]
// @Security     JWT
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	uc := usecase.NewGetUserUseCase(h.repo)
	u, err := uc.Execute(usecase.GetUserInputDTO{
		ID: id,
	})
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	slog.Info("User get", "users", u)
	h.tools.ResponseJSON(w, h.tools.NewResponse(u, http.StatusOK))

}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} tools.Response{data=[]usecase.GetUsersOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user [get]
// @Security     JWT
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, offset, err := h.tools.GetLimitOffsetFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	input := usecase.GetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	uc := usecase.NewGetUsersUseCase(h.repo)
	u, err := uc.Execute(input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	slog.Info("Users getting", "users", u)
	h.tools.ResponseJSON(w, h.tools.NewResponse(u, http.StatusOK))
}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body usecase.UpdateUserInputDTO true "User"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /user/{id} [put]
// @Security     JWT
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	var input usecase.UpdateUserInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	uc := usecase.NewUpdateUserUseCase(h.repo, id)
	err = uc.Execute(input)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err2.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	slog.Info("User updated")
	h.tools.ResponseJSON(w, h.tools.NewResponse(nil, http.StatusOK))
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
// @Security     JWT
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.tools.GetIDFromURL(r)
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	err = h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	uc := usecase.NewDeleteUserUseCase(h.repo)
	err = uc.Execute(usecase.DeleteUserInputDTO{
		ID: id,
	})
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err2.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		}
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		h.tools.ResponseErrorJSON(w, h.tools.NewResponseError(err.Error(), http.StatusBadRequest, "BAD_REQUEST"))
		return
	}
	slog.Info("User deleted")
	h.tools.ResponseJSON(w, h.tools.NewResponse(nil, http.StatusOK))
}
