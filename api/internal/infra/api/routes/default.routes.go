package routes

import (
	"fmt"
	"net/http"
	"strings"

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

func (route *Route) getTestHashValidate(router chi.Router) {
	router.Get("/hash/{hash}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hash")
		hash := chi.URLParam(r, "hash")
		resp, err := http.Post("http://localhost:3000/api/v1/auth/verify-user", "application/json", strings.NewReader(`{"hash":"`+hash+`"}`))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		w.Write([]byte("done"))
	})
}
