package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/handler"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/service"
)

func (route *Route) getAvatarRoutes(router chi.Router) {
	repo := repository.NewAvatarRepository(route.dbConn)
	service := service.NewAvatarService(repo)
	handler := handler.NewAvatarHandler(service)

	router.Route("/avatar", func(r chi.Router) {
		r.Get("/", handler.GetAvatars)
		r.Get("/{id}", handler.GetAvatar)
		r.Post("/", handler.CreateAvatar)
		r.Put("/{id}", handler.UpdateAvatar)
		r.Delete("/{id}", handler.DeleteAvatar)

	})
}
