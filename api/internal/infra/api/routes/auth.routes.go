package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getAuthRoutes(router chi.Router) {
	handler := handler.NewAuthHandler()
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.RefreshToken)
		r.Post("/register", handler.Register)
		r.Patch("/{id}/update-password", handler.UpdateUserPassword)
	})
}
