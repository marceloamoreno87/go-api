package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/infra/api"
	infraMiddleware "github.com/marceloamoreno/goapi/internal/infra/api/middleware"
	"github.com/marceloamoreno/goapi/internal/infra/database"
)

func Bootstrap() {

	r := startRouter()
	dbConn := startDbConn()
	startInfraMiddleware(r)
	startRoutes(r, dbConn)
	startServer(r)

}

func startServer(r *chi.Mux) {

	port := config.NewEnv().GetPort()
	slog.Info("Server started on port http://localhost:" + port + "/api/v1")
	slog.Info("Swagger started on port http://localhost:" + port + "/api/v1/swagger/index.html")
	slog.Info("Health started on port http://localhost:" + port + "/api/v1/health")

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}

}

func startRoutes(r *chi.Mux, dbConn *sql.DB) {

	api.NewRoutes(r, dbConn)
	slog.Info("Routes OK")

}

func startInfraMiddleware(r *chi.Mux) {
	infraMiddleware.NewLogMiddleware(r).LogMiddleware()
	slog.Info("Logger OK")

	infraMiddleware.NewCorsMiddleware(r).CorsMiddleware()
	slog.Info("Cors OK")
}

func startRouter() (r *chi.Mux) {
	r = chi.NewRouter()
	return
}

func startDbConn() (dbConn *sql.DB) {

	db := database.NewDatabase()

	err := db.SetDbConn()
	if err != nil {
		panic(err)
	}

	dbConn = db.GetDbConn()
	slog.Info("Database OK")

	return
}
