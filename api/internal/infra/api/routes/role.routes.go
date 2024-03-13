package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/handler"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
)

func (route *Route) getRoleRoutes(router chi.Router) {
	repo := repository.NewRoleRepository()
	service := service.NewRoleService(repo)
	handler := handler.NewRoleHandler(service)

	router.Route("/role", func(r chi.Router) {
		r.Get("/", handler.GetRoles)
		r.Get("/{id}", handler.GetRole)
		r.Post("/", handler.CreateRole)
		r.Put("/{id}", handler.UpdateRole)
		r.Delete("/{id}", handler.DeleteRole)
		route.getRolePermissionsRoutes(r)

	})
}
