package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
)

type UserHandler struct {
	repo repository.UserRepositoryInterface
}

func NewUserHandler(
	repo repository.UserRepositoryInterface,
) *UserHandler {
	return &UserHandler{
		repo: repo,
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

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewUserService(h.repo).CreateUser(r.Body)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
			return
		}
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("User created")
	// TODO: Response

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

	id := chi.URLParam(r, "id")

	output, err := service.NewUserService(h.repo).GetUser(id)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("User found")
	// TODO: Response

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
	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := service.NewUserService(h.repo).GetUsers(limit, offset)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Users found")
	// TODO: Response
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
	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewUserService(h.repo).UpdateUser(id, r.Body)

	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
			return
		}
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("User updated")
	// TODO: Response
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
	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewUserService(h.repo).DeleteUser(id)
	if err != nil {
		err2 := h.repo.Rollback()
		if err2 != nil {
			slog.Info("err", err2)
			// TODO: Response error
		}
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = h.repo.Commit()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("User deleted")
	// TODO: Response
}
