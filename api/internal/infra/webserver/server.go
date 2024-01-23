package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"

	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/internal/infra/database"
	"github.com/marceloamoreno/izimoney/internal/routes"
	"github.com/marceloamoreno/izimoney/tools"
)

func StartServer() {
	r := chi.NewRouter()
	loggerMiddleware(r)
	corsMiddleware(r)
	loadRoutes(r)
	fmt.Println("Server started on port http://localhost:" + config.Environment.Port + "/api/v1")
	fmt.Println("Swagger started on port http://localhost:" + config.Environment.Port + "/api/v1/swagger/index.html")
	fmt.Println("Health started on port http://localhost:" + config.Environment.Port + "/api/v1/health")
	http.ListenAndServe(":"+config.Environment.Port, r)
}

func loadRoutes(r *chi.Mux) {
	handlerTools := tools.NewHandlerTools()
	db := database.GetQueries()
	route := routes.NewRoute(r, handlerTools, db)
	route.Mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.GetAuthRoutes(r)
			route.GetRoute(r)
			route.GetSwaggerRoutes(r)
			route.GetHealthRoutes(r)
		})

		r.Group(func(r chi.Router) {
			authMiddleware(r)
			route.GetUserRoutes(r)
		})
	})

}

func loggerMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
}

func corsMiddleware(r *chi.Mux) {
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
}

func authMiddleware(r chi.Router) {
	r.Use(jwtauth.Verifier(config.TokenAuth))
	r.Use(jwtauth.Authenticator(config.TokenAuth))
}
