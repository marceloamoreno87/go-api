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
	corsMiddleware(r)
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
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
}