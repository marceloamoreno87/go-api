package routes

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/goapi/api/docs"
	"github.com/marceloamoreno/goapi/config"

	authMiddleware "github.com/marceloamoreno/goapi/internal/domain/auth/middleware"
)

type Route struct {
	mux    config.MuxInterface
	dbConn config.DatabaseInterface
	jwt    config.JWTAuthInterface
}

func NewRoutes(
	mux config.MuxInterface,
	dbConn config.DatabaseInterface,
	jwt config.JWTAuthInterface,
) {
	route := &Route{
		mux:    mux,
		dbConn: dbConn,
		jwt:    jwt,
	}
	route.mux.GetMux().Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.getAuthRoutes(r)
			route.getUserNonAuthRoutes(r)
			route.getRoute(r)
			route.getSwaggerRoutes(r)
			route.getHealthRoutes(r)
			route.getTestHashValidate(r)
		})

		r.Group(func(r chi.Router) {
			authMiddleware.NewMiddleware(r).AuthMiddleware(jwt.GetJwtAuth())
			route.getUserRoutes(r)
			route.getRoleRoutes(r)
			route.getPermissionRoutes(r)
			route.getAvatarRoutes(r)
		})
	})
}
