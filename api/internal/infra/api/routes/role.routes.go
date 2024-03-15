package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getRoleRoutes(router chi.Router) {
	DB := config.NewSqlc(config.NewDatabase())
	handler := handler.NewRoleHandler(DB)
	router.Route("/role", func(r chi.Router) {
		r.Get("/", handler.GetRoles)
		r.Get("/{id}", handler.GetRole)
		r.Post("/", handler.CreateRole)
		r.Put("/{id}", handler.UpdateRole)
		r.Delete("/{id}", handler.DeleteRole)
		route.getRolePermissionsRoutes(r)

	})
}
