package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
)

func (route *Route) getUserRoutes(router chi.Router) {
	userRepo := repository.NewUserRepository(route.dbConn)
	service := service.NewUserService(userRepo)
	handler := handler.NewUserHandler(service)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}

func (route *Route) getUserNonAuthRoutes(router chi.Router) {
	userRepo := repository.NewUserRepository(route.dbConn)
	service := service.NewUserService(userRepo)
	handler := handler.NewUserHandler(service)
	router.Route("/user", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/verify-user", handler.UserVerify)
		r.Post("/forgot-password", handler.ForgotPassword)
		r.Post("/update-password", handler.UpdatePasswordUser)
	})
}
