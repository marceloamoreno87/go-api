package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/marceloamoreno/izimoney/configs"
	_ "github.com/marceloamoreno/izimoney/docs"
	"github.com/marceloamoreno/izimoney/internal/domain/user/handler"
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

func (r *Route) GetUserRoutes() {
	repository := repository.NewUserRepository(database.Db())
	userHandler := handler.NewUserHandler(repository, r.HandlerTools)

	r.Route("/user", func(r chi.Router) {

		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate-token", userHandler.GetJWT)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})
	r.Route("/teste", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator(configs.TokenAuth))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Teste"))
		})
	})
}

func (r *Route) GetSwaggerRoutes() {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+configs.Environment.Port+"/swagger/doc.json"),
	))

}
