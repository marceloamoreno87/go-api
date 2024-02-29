package routes

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/goapi/api/docs"

	authMiddleware "github.com/marceloamoreno/goapi/internal/domain/auth/middleware"
)

type Route struct {
	mux    *chi.Mux
	dbConn *sql.DB
}

func NewRoutes(
	mux *chi.Mux,
	dbConn *sql.DB,
) {
	route := &Route{
		mux:    mux,
		dbConn: dbConn,
	}
	route.mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.getAuthRoutes()
			route.getRoute()
			route.getSwaggerRoutes()
			route.getHealthRoutes()
		})

		r.Group(func(r chi.Router) {
			authMiddleware.NewMiddleware(r).AuthMiddleware()
			slog.Info("Auth OK")
			route.getUserRoutes()
			route.getRoleRoutes()
			route.getPermissionRoutes()
		})
	})
}
