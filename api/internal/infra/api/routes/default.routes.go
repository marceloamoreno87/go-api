package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (route *Route) getSwaggerRoutes(router chi.Router) {
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.Environment.GetPort()+"/api/v1/swagger/doc.json"),
	))
}

func (route *Route) getRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
	})
}

func (route *Route) getHealthRoutes(router chi.Router) {
	router.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
