package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getAvatarRoutes(router chi.Router) {
	handler := handler.NewAvatarHandler()

	router.Route("/avatar", func(r chi.Router) {
		r.Get("/", handler.GetAvatars)
		r.Get("/{id}", handler.GetAvatar)
		r.Post("/", handler.CreateAvatar)
		r.Put("/{id}", handler.UpdateAvatar)
		r.Delete("/{id}", handler.DeleteAvatar)

	})
}
