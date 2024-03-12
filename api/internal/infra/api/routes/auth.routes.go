package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	"github.com/marceloamoreno/goapi/internal/domain/auth/repository"
	"github.com/marceloamoreno/goapi/internal/domain/auth/service"
)

func (route *Route) getAuthRoutes(router chi.Router) {
	repo := repository.NewAuthRepository(route.dbConn)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.RefreshToken)
	})
}
