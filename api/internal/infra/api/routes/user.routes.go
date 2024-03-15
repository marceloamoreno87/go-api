package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getUserRoutes(router chi.Router) {
	DB := config.NewSqlc(config.NewDatabase())
	handler := handler.NewUserHandler(DB)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUserById)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}
