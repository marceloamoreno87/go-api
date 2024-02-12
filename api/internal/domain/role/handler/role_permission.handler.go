package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	PermissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type RolePermissionHandler struct {
	HandlerTools         api.HandlerToolsInterface
	RoleRepository       repository.RoleRepositoryInterface
	PermissionRepository PermissionRepository.PermissionRepositoryInterface
}

func NewRolePermissionHandler(
	roleRepository repository.RoleRepositoryInterface,
	permissionRepository PermissionRepository.PermissionRepositoryInterface,
	handlerTools api.HandlerToolsInterface,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
		HandlerTools:         handlerTools,
	}
}

// CreateRolePermission godoc
// @Summary Create Role Permission
// @Description Create Role Permission
// @Tags RolePermission
// @Accept  json
// @Produce  json
// @Param user body usecase.RolePermissionInputDTO true "RolePermission"
// @Success 200 {object} api.Response{data=entity.RolePermission}
// @Failure 400 {object} api.ResponseError{err=string}
// @Router /role/permission [post]
// @Security     JWT
func (h *RolePermissionHandler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {

	var rolePermission usecase.CreateRolePermissionInputDTO
	err := json.NewDecoder(r.Body).Decode(&rolePermission)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}

	uc := usecase.NewCreateRolePermissionUseCase(h.RoleRepository, h.PermissionRepository)
	u, err := uc.Execute(rolePermission)
	if err != nil {
		slog.Info("err", err)
		h.HandlerTools.ResponseErrorJSON(w, api.NewResponseErrorDefault(err.Error()))
		return
	}
	slog.Info("Role permission created", "Role permission", u)
	h.HandlerTools.ResponseJSON(w, u)
}
