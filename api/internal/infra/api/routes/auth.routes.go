package routes

import (
	"github.com/go-chi/chi/v5"

	authHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	authService "github.com/marceloamoreno/goapi/internal/domain/auth/service"
	authRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

func (r *Route) getAuthRoutes() {
	repo := authRepository.NewUserRepository(r.dbConn)
	service := authService.NewAuthService(repo)
	handler := authHandler.NewAuthHandler(service)
	r.mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
	})
}
