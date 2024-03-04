package routes

import (
	"github.com/go-chi/chi/v5"
	rolePermissionHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	rolePermissionRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	rolePermissionService "github.com/marceloamoreno/goapi/internal/domain/role/service"
)

func (route *Route) getRolePermissionsRoutes(router chi.Router) {
	repo := rolePermissionRepository.NewRolePermissionRepository(route.dbConn)
	service := rolePermissionService.NewRolePermissionService(repo)
	handler := rolePermissionHandler.NewRolePermissionHandler(service)

	router.Route("/{id}/permission", func(r chi.Router) {
		r.Get("/", handler.GetRolePermissions)
		r.Post("/", handler.CreateRolePermission)
		r.Put("/", handler.UpdateRolePermission)
	})

}
