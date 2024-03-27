package routes

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/goapi/api/docs"
	"github.com/marceloamoreno/goapi/config"
	// AuthMiddleware "github.com/marceloamoreno/goapi/internal/domain/user/middleware"
)

type Route struct {
	mux config.MuxInterface
	jwt config.JWTAuthInterface
}

func NewRoutes(
	mux config.MuxInterface,
	jwt config.JWTAuthInterface,
) {
	route := &Route{
		mux: mux,
		jwt: jwt,
	}
	route.mux.GetMux().Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.getAuthRoutes(r)
			route.getRoute(r)
			route.getSwaggerRoutes(r)
			route.getHealthRoutes(r)
			route.getTestUpdatePassword(r)
			route.getTestVerifyUser(r)

		})
		r.Group(func(r chi.Router) {
			// AuthMiddleware.NewMiddleware(r).AuthMiddleware(jwt.GetJwtAuth())
			route.getUserRoutes(r)
			route.getRoleRoutes(r)
			route.getPermissionRoutes(r)
			route.getAvatarRoutes(r)
		})
	})
}
