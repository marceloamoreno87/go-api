package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/marceloamoreno/goapi/config"
	authMiddleware "github.com/marceloamoreno/goapi/internal/domain/auth/middleware"
	"github.com/marceloamoreno/goapi/internal/infra/database"
	infraMiddleware "github.com/marceloamoreno/goapi/internal/infra/webserver/middleware"
)

func StartServer() {
	r := chi.NewRouter()

	infraMiddleware.NewLogMiddleware(r).LogMiddleware()
	infraMiddleware.NewCorsMiddleware(r).CorsMiddleware()

	db := database.NewDatabase()
	err := db.SetDbConn()
	if err != nil {
		panic(err)
	}
	dbConn := db.GetDbConn()

	slog.Info("Database OK")

	loadRoutes(r, dbConn)

	port := config.NewEnv().GetPort()
	slog.Info("Server started on port http://localhost:" + port + "/api/v1")
	slog.Info("Swagger started on port http://localhost:" + port + "/api/v1/swagger/index.html")
	slog.Info("Health started on port http://localhost:" + port + "/api/v1/health")

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}

func loadRoutes(
	r *chi.Mux,
	dbConn *sql.DB,
) {
	route := NewRoute(r, dbConn)
	route.mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.GetAuthRoutes(r)
			route.GetRoute(r)
			route.GetSwaggerRoutes(r)
			route.GetHealthRoutes(r)
		})

		r.Group(func(r chi.Router) {
			authMiddleware.NewMiddleware(r).AuthMiddleware()
			route.GetUserRoutes(r)
			route.GetRoleRoutes(r)
			route.GetPermissionRoutes(r)
		})
	})
	slog.Info("Routes OK")
}
