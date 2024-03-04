package routes

import (
	"github.com/go-chi/chi/v5"
	userHandler "github.com/marceloamoreno/goapi/internal/domain/user/handler"
	userRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
	userService "github.com/marceloamoreno/goapi/internal/domain/user/service"
)

func (route *Route) getUserRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(route.dbConn)
	service := userService.NewUserService(repo)
	handler := userHandler.NewUserHandler(service)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}
