package routes

import (
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

func (route *Route) getTestVerifyUser(router chi.Router) {
	router.Get("/verify-user/hash/{hash}", func(w http.ResponseWriter, r *http.Request) {

		// Test verify user
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

func (route *Route) getTestUpdatePassword(router chi.Router) {
	router.Get("/update-password/hash/{hash}", func(w http.ResponseWriter, r *http.Request) {
		// Test forgot password
		hash := chi.URLParam(r, "hash")
		resp, err := http.Post("http://localhost:3000/api/v1/auth/update-password", "application/json", strings.NewReader(`{"hash":"`+hash+`","password":"123456"}`))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		w.Write([]byte("done"))

	})
}
