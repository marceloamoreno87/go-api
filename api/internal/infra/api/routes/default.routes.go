package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (r *Route) getSwaggerRoutes() {
	r.mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.NewEnv().GetPort()+"/api/v1/swagger/doc.json"),
	))
}

func (r *Route) getRoute() {
	r.mux.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
	})
}

func (r *Route) getHealthRoutes() {
	r.mux.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
