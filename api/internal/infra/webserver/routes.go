package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	_ "github.com/marceloamoreno/goapi/docs"
	AuthHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	UserHandler "github.com/marceloamoreno/goapi/internal/domain/user/handler"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/infra/database"
	"github.com/marceloamoreno/goapi/pkg/api"
	HttpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	HandlerTools *api.HandlerTools
	Mux          *chi.Mux
	DBConn       *sql.DB
}

func NewRoute(r *chi.Mux, handlerTools *api.HandlerTools) *Route {
	DBConn, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}
	return &Route{
		Mux:          r,
		HandlerTools: handlerTools,
		DBConn:       DBConn,
	}
}

func (r *Route) GetAuthRoutes(router chi.Router) {
	AuthRepository := repository.NewUserRepository(r.DBConn)
	AuthHandler := AuthHandler.NewAuthHandler(AuthRepository, r.HandlerTools)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", AuthHandler.GetJWT)
		r.Post("/token/refresh", AuthHandler.GetRefreshJWT)
	})
}

func (r *Route) GetUserRoutes(router chi.Router) {
	UserRepository := repository.NewUserRepository(r.DBConn)
	UserHandler := UserHandler.NewUserHandler(UserRepository, r.HandlerTools)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", UserHandler.GetUsers)
		r.Get("/{id}", UserHandler.GetUser)
		r.Post("/", UserHandler.CreateUser)
		r.Put("/{id}", UserHandler.UpdateUser)
		r.Delete("/{id}", UserHandler.DeleteUser)
	})
}

func (r *Route) GetSwaggerRoutes(router chi.Router) {
	router.Get("/swagger/*", HttpSwagger.Handler(
		HttpSwagger.URL("http://localhost:"+config.Environment.Port+"/api/v1/swagger/doc.json"),
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
