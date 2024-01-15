package configs

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marceloamoreno/izimoney/internal/routes"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	loadRoutes(r)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func loadRoutes(r *chi.Mux) {
	queries := Queries()
	routes.NewRoute(r, queries).GetUserRoutes()
	routes.NewRoute(r, queries).GetSwaggerRoutes()

}
