package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/marceloamoreno/goapi/config"
	AuthMiddleware "github.com/marceloamoreno/goapi/internal/domain/auth/middleware"
	CorsMiddleware "github.com/marceloamoreno/goapi/internal/infra/webserver/middleware"
	LogMiddleware "github.com/marceloamoreno/goapi/internal/infra/webserver/middleware"
	"github.com/marceloamoreno/goapi/pkg/api"
)

func StartServer() {
	r := chi.NewRouter()
	LogMiddleware.NewLogMiddleware(r).LogMiddleware()
	CorsMiddleware.NewCorsMiddleware(r).CorsMiddleware()
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
	handlerTools := api.NewHandlerTools()
	route := NewRoute(r, handlerTools)
	route.Mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.GetAuthRoutes(r)
			route.GetRoute(r)
			route.GetSwaggerRoutes(r)
			route.GetHealthRoutes(r)
		})

		r.Group(func(r chi.Router) {
			AuthMiddleware.NewMiddleware(r).AuthMiddleware()
			route.GetUserRoutes(r)
			route.GetRoleRoutes(r)
			route.GetPermissionRoutes(r)
			route.GetRolePermissionRoutes(r)
		})
	})
	slog.Info("Routes OK")
}
