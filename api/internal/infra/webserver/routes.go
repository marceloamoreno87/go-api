package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	_ "github.com/marceloamoreno/goapi/docs"
	authHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	permissionHandler "github.com/marceloamoreno/goapi/internal/domain/permission/handler"
	permissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	roleHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	roleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	userHandler "github.com/marceloamoreno/goapi/internal/domain/user/handler"
	userRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/pkg/api"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	handlerTools *api.HandlerTools
	mux          *chi.Mux
	dbConn       *sql.DB
}

func NewRoute(
	r *chi.Mux,
	handlerTools *api.HandlerTools,
	db *sql.DB,
) *Route {
	return &Route{
		mux:          r,
		handlerTools: handlerTools,
		dbConn:       db,
	}
}

func (r *Route) GetAuthRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(r.dbConn)
	handler := authHandler.NewAuthHandler(repo, r.handlerTools)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", handler.GetJWT)
		r.Post("/token/refresh", handler.GetRefreshJWT)
	})
}

func (r *Route) GetUserRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(r.dbConn)
	handler := userHandler.NewUserHandler(repo, r.handlerTools)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}

func (r *Route) GetRoleRoutes(router chi.Router) {
	repo := roleRepository.NewRoleRepository(r.dbConn)
	repo2 := roleRepository.NewRolePermissionRepository(r.dbConn)
	handler := roleHandler.NewRoleHandler(repo, r.handlerTools)
	handler2 := roleHandler.NewRolePermissionHandler(repo2, r.handlerTools)
	router.Route("/role", func(r chi.Router) {
		r.Get("/", handler.GetRoles)
		r.Get("/{id}", handler.GetRole)
		r.Post("/", handler.CreateRole)
		r.Put("/{id}", handler.UpdateRole)
		r.Delete("/{id}", handler.DeleteRole)

		r.Route("/{id}/permission", func(r chi.Router) {
			r.Get("/", handler2.GetRolePermissions)
			r.Post("/", handler2.CreateRolePermission)
			r.Put("/", handler2.UpdateRolePermission)
		})

	})
}

func (r *Route) GetPermissionRoutes(router chi.Router) {
	repo := permissionRepository.NewPermissionRepository(r.dbConn)
	handler := permissionHandler.NewPermissionHandler(repo, r.handlerTools)
	router.Route("/permission", func(r chi.Router) {
		r.Get("/", handler.GetPermissions)
		r.Get("/{id}", handler.GetPermission)
		r.Post("/", handler.CreatePermission)
		r.Put("/{id}", handler.UpdatePermission)
		r.Delete("/{id}", handler.DeletePermission)

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
