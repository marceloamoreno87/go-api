package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/user/handler"
)

func (route *Route) getUserRoutes(router chi.Router) {
	handler := handler.NewUserHandler()
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
		r.Post("/update-password", handler.UpdateUserPassword)
		r.Post("/verify-user", handler.VerifyUser)
		r.Post("/forgot-password", handler.ForgotPassword)
	})
}
