package routes

import (
	"github.com/go-chi/chi/v5"
	avatarHandler "github.com/marceloamoreno/goapi/internal/domain/avatar/handler"
	avatarRepository "github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
	avatarService "github.com/marceloamoreno/goapi/internal/domain/avatar/service"
)

func (r *Route) getAvatarRoutes(router chi.Router) {
	repo := avatarRepository.NewAvatarRepository(r.dbConn)
	service := avatarService.NewAvatarService(repo)
	handler := avatarHandler.NewAvatarHandler(service)

	router.Route("/avatar", func(r chi.Router) {
		r.Get("/", handler.GetAvatars)
		r.Get("/{id}", handler.GetAvatar)
		r.Post("/", handler.CreateAvatar)
		r.Put("/{id}", handler.UpdateAvatar)
		r.Delete("/{id}", handler.DeleteAvatar)

	})
}
