package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/internal/routes"
	"github.com/marceloamoreno/izimoney/tools"
)

func StartServer() {
	r := chi.NewRouter()
	loggerMiddleware(r)
	corsMiddleware(r)
	loadRoutes(r)
	http.ListenAndServe(":"+config.Environment.Port, r)
}

func loadRoutes(r *chi.Mux) {
	handlerTools := tools.NewHandlerTools()
	route := routes.NewRoute(r, handlerTools)
	route.GetUserRoutes()
	route.GetSwaggerRoutes()

	// Example of route with JWT
	route.GetExampleRoute()

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
