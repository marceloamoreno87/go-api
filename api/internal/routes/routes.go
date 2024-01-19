package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/izimoney/config"
	_ "github.com/marceloamoreno/izimoney/docs"
	authHandler "github.com/marceloamoreno/izimoney/internal/domain/auth/handler"
	userHandler "github.com/marceloamoreno/izimoney/internal/domain/user/handler"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
	"github.com/marceloamoreno/izimoney/internal/infra/database"
	"github.com/marceloamoreno/izimoney/tools"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	*tools.HandlerTools
	*chi.Mux
}

func NewRoute(r *chi.Mux, handlerTools *tools.HandlerTools) *Route {
	return &Route{
		Mux:          r,
		HandlerTools: handlerTools,
	}
}

func (r *Route) GetAuthRoutes(router chi.Router) {
	repository := repository.NewUserRepository(database.Db())
	authHandler := authHandler.NewAuthHandler(repository, r.HandlerTools)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", authHandler.GetJWT)
	})
}

func (r *Route) GetUserRoutes(router chi.Router) {
	repository := repository.NewUserRepository(database.Db())
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
		httpSwagger.URL("http://localhost:"+config.Environment.Port+"/swagger/doc.json"),
	))
}

// Example of route with JWT
func (r *Route) GetExampleRoute(router chi.Router) {
	router.Route("/teste", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Teste"))
		})
	})
}

func (r *Route) GetHealthRoutes(router chi.Router) {
	router.Route("/api/v1/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
