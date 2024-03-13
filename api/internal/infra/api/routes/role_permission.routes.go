package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/role/handler"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/service"
)

func (route *Route) getRolePermissionsRoutes(router chi.Router) {
	repo := repository.NewRolePermissionRepository()
	service := service.NewRolePermissionService(repo)
	handler := handler.NewRolePermissionHandler(service)

	router.Route("/{id}/permission", func(r chi.Router) {
		r.Get("/", handler.GetRolePermissions)
		r.Post("/", handler.CreateRolePermission)
		r.Put("/", handler.UpdateRolePermission)
	})

}
