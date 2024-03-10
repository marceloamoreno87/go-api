package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
)

func (route *Route) getAuthRoutes(router chi.Router) {
	userRepo := repository.NewUserRepository(route.dbConn)
	service := service.NewUserService(userRepo)
	handler := handler.NewAuthHandler(service)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
		r.Post("/verify-user", handler.UserVerify)
		r.Post("/forgot-password", handler.ForgotPassword)
		r.Post("/update-password", handler.UpdatePasswordUser)
	})
}
