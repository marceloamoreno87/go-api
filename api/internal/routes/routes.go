package routes

import (
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/izimoney/docs"
	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
	"github.com/marceloamoreno/izimoney/tools"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	*tools.HandlerTools
	*chi.Mux
	*db.Queries
}

func NewRoute(r *chi.Mux, db *db.Queries, t *tools.HandlerTools) *Route {
	return &Route{
		tools.NewHandlerTools(),
		r,
		db,
	}
}

func (r *Route) GetUserRoutes() {
	UserHandler := handler.NewUserHandler(r.Queries, r.HandlerTools)
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
		httpSwagger.URL("http://localhost:"+os.Getenv("PORT")+"/swagger/doc.json"),
	))
}
