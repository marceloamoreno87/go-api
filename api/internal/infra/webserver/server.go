package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/marceloamoreno/goapi/config"
	infraMiddleware "github.com/marceloamoreno/goapi/internal/infra/api/middleware"
	"github.com/marceloamoreno/goapi/internal/infra/api/routes"
	"github.com/marceloamoreno/goapi/internal/infra/database"
)

func Bootstrap() {
	mux := startRouter()
	dbConn := startDbConn()
	startInfraMiddleware(mux)
	startRoutes(mux, dbConn)
	startServer(mux)
}

func startServer(mux *chi.Mux) {
	port := config.NewEnv().GetPort()
	slog.Info("Server started on port http://localhost:" + port + "/api/v1")
	slog.Info("Swagger started on port http://localhost:" + port + "/api/v1/swagger/index.html")
	slog.Info("Health started on port http://localhost:" + port + "/api/v1/health")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

func startRoutes(mux *chi.Mux, dbConn *sql.DB) {
	routes.NewRoutes(mux, dbConn)
	slog.Info("Routes OK")
}

func startInfraMiddleware(mux *chi.Mux) {
	infraMiddleware.NewLogMiddleware(mux).LogMiddleware()
	slog.Info("Logger OK")
	infraMiddleware.NewCorsMiddleware(mux).CorsMiddleware()
	slog.Info("Cors OK")
}

func startRouter() (mux *chi.Mux) {
	mux = chi.NewRouter()
	return
}

func startDbConn() (dbConn *sql.DB) {
	db := database.NewDatabase()
	if err := db.SetDbConn(); err != nil {
		panic(err)
	}
	dbConn = db.GetDbConn()
	slog.Info("Database OK")
	return
}
