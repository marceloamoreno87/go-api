package routes

import (
	"github.com/go-chi/chi/v5"
	roleHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	roleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	roleService "github.com/marceloamoreno/goapi/internal/domain/role/service"

	rolePermissionHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	rolePermissionRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	rolePermissionService "github.com/marceloamoreno/goapi/internal/domain/role/service"
)

func (r *Route) getRoleRoutes(router chi.Router) {
	repo := roleRepository.NewRoleRepository(r.dbConn)
	service := roleService.NewRoleService(repo)
	handler := roleHandler.NewRoleHandler(service)

	repo2 := rolePermissionRepository.NewRolePermissionRepository(r.dbConn)
	service2 := rolePermissionService.NewRolePermissionService(repo2)
	handler2 := rolePermissionHandler.NewRolePermissionHandler(service2)

	router.Route("/role", func(r chi.Router) {
		r.Get("/", handler.GetRoles)
		r.Get("/{id}", handler.GetRole)
		r.Post("/", handler.CreateRole)
		r.Put("/{id}", handler.UpdateRole)
		r.Delete("/{id}", handler.DeleteRole)

		r.Route("/{id}/permission", func(r chi.Router) {
			r.Get("/", handler2.GetRolePermissions)
			r.Post("/", handler2.CreateRolePermission)
			r.Put("/", handler2.UpdateRolePermission)
		})

	})
}
