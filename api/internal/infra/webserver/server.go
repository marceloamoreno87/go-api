package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/exp/slog"

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
	slog.Info("Server started on port http://localhost:" + config.Environment.Port + "/api/v1")
	slog.Info("Swagger started on port http://localhost:" + config.Environment.Port + "/api/v1/swagger/index.html")
	slog.Info("Health started on port http://localhost:" + config.Environment.Port + "/api/v1/health")
	err := http.ListenAndServe(":"+config.Environment.Port, r)
	if err != nil {
		panic(err)
	}

}

func loadRoutes(r *chi.Mux) {
	handlerTools := tools.NewHandlerTools()
	db := database.GetQueries()
	slog.Info("Database OK")
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
	slog.Info("Routes OK")
}

func loggerMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	slog.Info("Logger OK")
}

func corsMiddleware(r *chi.Mux) {
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
	slog.Info("Cors OK")
}

func authMiddleware(r chi.Router) {
	r.Use(jwtauth.Verifier(config.TokenAuth))
	r.Use(jwtauth.Authenticator(config.TokenAuth))
	slog.Info("Auth OK")
}
