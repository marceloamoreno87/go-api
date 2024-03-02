package routes

import (
	"github.com/go-chi/chi/v5"
	permissionHandler "github.com/marceloamoreno/goapi/internal/domain/permission/handler"
	permissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	permissionService "github.com/marceloamoreno/goapi/internal/domain/permission/service"
)

func (r *Route) getPermissionRoutes(router chi.Router) {
	repo := permissionRepository.NewPermissionRepository(r.dbConn)
	service := permissionService.NewPermissionService(repo)
	handler := permissionHandler.NewPermissionHandler(service)

	router.Route("/permission", func(r chi.Router) {
		r.Get("/", handler.GetPermissions)
		r.Get("/{id}", handler.GetPermission)
		r.Post("/", handler.CreatePermission)
		r.Put("/{id}", handler.UpdatePermission)
		r.Delete("/{id}", handler.DeletePermission)

	})
}
