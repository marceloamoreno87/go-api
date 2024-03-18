package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getAuthRoutes(router chi.Router) {
	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()
	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.RegisterUser)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.RefreshToken)
	})
}
