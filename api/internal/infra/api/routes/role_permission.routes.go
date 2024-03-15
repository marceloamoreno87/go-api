package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getRolePermissionsRoutes(router chi.Router) {
	handler := handler.NewRolePermissionHandler()

	router.Route("/{id}/permission", func(r chi.Router) {
		r.Get("/", handler.GetRolePermissions)
		r.Post("/", handler.CreateRolePermission)
		r.Delete("/", handler.DeleteRolePermission)
	})

}
