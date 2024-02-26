package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
)

type RoleHandler struct {
	repo repository.RoleRepositoryInterface
}

func NewRoleHandler(
	repo repository.RoleRepositoryInterface,
) *RoleHandler {
	return &RoleHandler{
		repo: repo,
	}
}

// GetRole godoc
// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} tools.Response{data=usecase.GetRoleOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role/{id} [get]
// @Security     JWT
func (h *RoleHandler) GetRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	output, err := service.NewRoleService(h.repo).GetRole(id)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Role found")
	// TODO: Response
}

// GetRoles godoc
// @Summary Get Roles
// @Description Get Roles
// @Tags Role
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} tools.Response{data=[]usecase.GetRolesOutputDTO}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role [get]
// @Security     JWT
func (h *RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request) {

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := service.NewRoleService(h.repo).GetRoles(limit, offset)
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	slog.Info("Roles found")
	// TODO: Response

}

// CreateRole godoc
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateRoleInputDTO true "Role"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role [post]
// @Security     JWT
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = service.NewRoleService(h.repo).CreateRole(r.Body)
	if err != nil {
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

	slog.Info("Role created")
	// TODO: Response
}

// UpdateRole godoc
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param role body usecase.UpdateUserInputDTO true "Role"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Router /role/{id} [put]
// @Security     JWT
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}
	err = service.NewRoleService(h.repo).UpdateRole(id, r.Body)
	if err != nil {
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

	slog.Info("Role updated")
	// TODO: Response

}

// DeleteRole godoc
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} tools.Response{data=nil}
// @Failure 400 {object} tools.ResponseError{err=string}
// @Security ApiKeyAuth
// @Router /role/{id} [delete]
// @Security     JWT
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.repo.Begin()
	if err != nil {
		slog.Info("err", err)
		// TODO: Response error
		return
	}

	err = service.NewRoleService(h.repo).DeleteRole(id)
	if err != nil {
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

	slog.Info("Role deleted")
	// TODO: Response

}
