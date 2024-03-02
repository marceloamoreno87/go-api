package routes

import (
	"github.com/go-chi/chi/v5"

	authHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	userRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
	userService "github.com/marceloamoreno/goapi/internal/domain/user/service"
)

func (r *Route) getAuthRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(r.dbConn)
	service := userService.NewUserService(repo)
	handler := authHandler.NewAuthHandler(service)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
	})
}
