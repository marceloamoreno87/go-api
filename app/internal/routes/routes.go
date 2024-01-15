package routes

import (
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/izimoney/docs"
	"github.com/marceloamoreno/izimoney/internal/db"
	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	*chi.Mux
	*db.Queries
}

func NewRoute(r *chi.Mux, db *db.Queries) *Route {
	return &Route{
		r,
		db,
	}
}

func (r *Route) GetUserRoutes() {
	UserHandler := handler.NewUserHandler(r.Queries)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Put("/", UserHandler.CreateUser)
		r.Patch("/", handler.UpdateUser)
		r.Delete("/", handler.DeleteUser)
	})
}

func (r *Route) GetSwaggerRoutes() {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+os.Getenv("PORT")+"/swagger/doc.json"),
	))
}
