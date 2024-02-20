package webserver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/config"
	_ "github.com/marceloamoreno/goapi/docs"
	AuthHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	PermissionHandler "github.com/marceloamoreno/goapi/internal/domain/permission/handler"
	PermissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	RoleHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	RolePermissionHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	RolePermissionRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	UserHandler "github.com/marceloamoreno/goapi/internal/domain/user/handler"
	UserRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/pkg/api"
	HttpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	HandlerTools *api.HandlerTools
	Mux          *chi.Mux
	DBConn       *sql.DB
	Transaction  api.TransactionInterface
}

func NewRoute(r *chi.Mux, handlerTools *api.HandlerTools, db *sql.DB) *Route {
	return &Route{
		Mux:          r,
		HandlerTools: handlerTools,
		Transaction:  api.NewTransaction(db),
		DBConn:       db,
	}
}

func (r *Route) GetAuthRoutes(router chi.Router) {
	AuthRepository := UserRepository.NewUserRepository(r.DBConn)
	AuthHandler := AuthHandler.NewAuthHandler(AuthRepository, r.HandlerTools)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", AuthHandler.GetJWT)
		r.Post("/token/refresh", AuthHandler.GetRefreshJWT)
	})
}

func (r *Route) GetUserRoutes(router chi.Router) {
	userRepository := UserRepository.NewUserRepository(r.DBConn)
	userHandler := UserHandler.NewUserHandler(userRepository, r.HandlerTools, r.Transaction)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})
}

func (r *Route) GetRoleRoutes(router chi.Router) {
	RoleRepository := RoleRepository.NewRoleRepository(r.DBConn)
	RoleHandler := RoleHandler.NewRoleHandler(RoleRepository, r.HandlerTools)
	RolePermissionRepository := RolePermissionRepository.NewRolePermissionRepository(r.DBConn)
	RolePermissionHandler := RolePermissionHandler.NewRolePermissionHandler(RolePermissionRepository, r.HandlerTools)
	router.Route("/role", func(r chi.Router) {
		r.Get("/", RoleHandler.GetRoles)
		r.Get("/{id}", RoleHandler.GetRole)
		r.Post("/", RoleHandler.CreateRole)
		r.Put("/{id}", RoleHandler.UpdateRole)
		r.Delete("/{id}", RoleHandler.DeleteRole)

		r.Route("/{id}/permission", func(r chi.Router) {
			r.Get("/", RolePermissionHandler.GetRolePermissions)
			r.Post("/", RolePermissionHandler.CreateRolePermission)
			r.Put("/", RolePermissionHandler.UpdateRolePermission)
		})

	})
}

func (r *Route) GetPermissionRoutes(router chi.Router) {
	PermissionRepository := PermissionRepository.NewPermissionRepository(r.DBConn)
	PermissionHandler := PermissionHandler.NewPermissionHandler(PermissionRepository, r.HandlerTools)
	router.Route("/permission", func(r chi.Router) {
		r.Get("/", PermissionHandler.GetPermissions)
		r.Get("/{id}", PermissionHandler.GetPermission)
		r.Post("/", PermissionHandler.CreatePermission)
		r.Put("/{id}", PermissionHandler.UpdatePermission)
		r.Delete("/{id}", PermissionHandler.DeletePermission)

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
