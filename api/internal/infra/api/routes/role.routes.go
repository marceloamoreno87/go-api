package routes

import (
	"github.com/go-chi/chi/v5"
	roleHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	roleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	roleService "github.com/marceloamoreno/goapi/internal/domain/role/service"
)

func (route *Route) getRoleRoutes(router chi.Router) {
	repo := roleRepository.NewRoleRepository(route.dbConn)
	service := roleService.NewRoleService(repo)
	handler := roleHandler.NewRoleHandler(service)

	router.Route("/role", func(r chi.Router) {
		r.Get("/", handler.GetRoles)
		r.Get("/{id}", handler.GetRole)
		r.Post("/", handler.CreateRole)
		r.Put("/{id}", handler.UpdateRole)
		r.Delete("/{id}", handler.DeleteRole)
		route.getRolePermissionsRoutes(r)

	})
}
