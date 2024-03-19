package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getPermissionRoutes(router chi.Router) {
	handler := handler.NewPermissionHandler()
	router.Route("/permission", func(r chi.Router) {
		r.Get("/", handler.GetPermissions)
		r.Get("/{id}", handler.GetPermission)
		r.Post("/", handler.CreatePermission)
		r.Put("/{id}", handler.UpdatePermission)
		r.Delete("/{id}", handler.DeletePermission)
	})
}
