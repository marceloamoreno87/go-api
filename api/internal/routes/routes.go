package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/izimoney/configs"
	_ "github.com/marceloamoreno/izimoney/docs"
	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/internal/infra/database"
	"github.com/marceloamoreno/izimoney/tools"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	*tools.HandlerTools
	*chi.Mux
}

func NewRoute(r *chi.Mux, handlerTools *tools.HandlerTools) *Route {
	return &Route{
		Mux:          r,
		HandlerTools: handlerTools,
	}
}

func (r *Route) GetUserRoutes() {
	repository := repository.NewUserRepository(database.Db())
	UserHandler := handler.NewUserHandler(repository, r.HandlerTools)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", UserHandler.GetUsers)
		r.Get("/{id}", UserHandler.GetUser)
		r.Post("/", UserHandler.CreateUser)
		r.Put("/{id}", UserHandler.UpdateUser)
		r.Delete("/{id}", UserHandler.DeleteUser)
	})
}

func (r *Route) GetSwaggerRoutes() {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+configs.Environment.Port+"/swagger/doc.json"),
	))

}
