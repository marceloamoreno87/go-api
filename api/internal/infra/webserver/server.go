package webserver

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marceloamoreno/izimoney/internal/infra/database"
	"github.com/marceloamoreno/izimoney/internal/routes"
	"github.com/marceloamoreno/izimoney/tools"
)

func StartServer() {
	r := chi.NewRouter()
	loggerMiddleware(r)
	loadRoutes(r)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func loadRoutes(r *chi.Mux) {
	HandlerTools := tools.NewHandlerTools()
	route := routes.NewRoute(r, database.Db(), HandlerTools)
	route.GetUserRoutes()
	route.GetSwaggerRoutes()
}

func loggerMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
}

func corsMiddleware(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
