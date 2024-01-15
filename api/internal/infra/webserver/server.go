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
	r.Use(middleware.Logger)
	loadRoutes(r)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func loadRoutes(r *chi.Mux) {
	HandlerTools := tools.NewHandlerTools()
	route := routes.NewRoute(r, database.Db(), HandlerTools)
	route.GetUserRoutes()
	route.GetSwaggerRoutes()

}
