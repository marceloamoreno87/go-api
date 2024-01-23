package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/izimoney/config"
	_ "github.com/marceloamoreno/izimoney/docs"
	authHandler "github.com/marceloamoreno/izimoney/internal/domain/auth/handler"
	userHandler "github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
	"github.com/marceloamoreno/izimoney/tools"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	HandlerTools *tools.HandlerTools
	Mux          *chi.Mux
	DB           *db.Queries
}

func NewRoute(r *chi.Mux, handlerTools *tools.HandlerTools, db *db.Queries) *Route {
	return &Route{
		Mux:          r,
		HandlerTools: handlerTools,
		DB:           db,
	}
}

func (r *Route) GetAuthRoutes(router chi.Router) {
	repository := repository.NewUserRepository(r.DB)
	authHandler := authHandler.NewAuthHandler(repository, r.HandlerTools)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", authHandler.GetJWT)
	})
}

func (r *Route) GetUserRoutes(router chi.Router) {
	repository := repository.NewUserRepository(r.DB)
	userHandler := userHandler.NewUserHandler(repository, r.HandlerTools)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})
}

func (r *Route) GetSwaggerRoutes(router chi.Router) {
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.Environment.Port+"/api/v1/swagger/doc.json"),
	))
}

func (r *Route) GetRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
	})
}

func (r *Route) GetHealthRoutes(router chi.Router) {
	router.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
