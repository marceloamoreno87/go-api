package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/permission/handler"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/service"
)

func (route *Route) getPermissionRoutes(router chi.Router) {
	repo := repository.NewPermissionRepository()
	service := service.NewPermissionService(repo)
	handler := handler.NewPermissionHandler(service)

	router.Route("/permission", func(r chi.Router) {
		r.Get("/", handler.GetPermissions)
		r.Get("/{id}", handler.GetPermission)
		r.Post("/", handler.CreatePermission)
		r.Put("/{id}", handler.UpdatePermission)
		r.Delete("/{id}", handler.DeletePermission)

	})
}
