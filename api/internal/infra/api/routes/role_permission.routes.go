package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getRolePermissionsRoutes(router chi.Router) {
	DB := config.NewSqlc(config.NewDatabase())
	handler := handler.NewRolePermissionHandler(DB)

	router.Route("/{id}/permission", func(r chi.Router) {
		r.Get("/", handler.GetRolePermissions)
		r.Post("/", handler.CreateRolePermission)
		r.Delete("/", handler.DeleteRolePermission)
	})

}
