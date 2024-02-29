package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/marceloamoreno/goapi/api/docs"
	"github.com/marceloamoreno/goapi/config"

	authHandler "github.com/marceloamoreno/goapi/internal/domain/auth/handler"
	authMiddleware "github.com/marceloamoreno/goapi/internal/domain/auth/middleware"
	authService "github.com/marceloamoreno/goapi/internal/domain/auth/service"

	roleHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	roleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	roleService "github.com/marceloamoreno/goapi/internal/domain/role/service"

	permissionHandler "github.com/marceloamoreno/goapi/internal/domain/permission/handler"
	permissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	permissionService "github.com/marceloamoreno/goapi/internal/domain/permission/service"

	rolePermissionHandler "github.com/marceloamoreno/goapi/internal/domain/role/handler"
	rolePermissionRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
	rolePermissionService "github.com/marceloamoreno/goapi/internal/domain/role/service"

	userHandler "github.com/marceloamoreno/goapi/internal/domain/user/handler"
	userRepository "github.com/marceloamoreno/goapi/internal/domain/user/repository"
	userService "github.com/marceloamoreno/goapi/internal/domain/user/service"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	mux    *chi.Mux
	dbConn *sql.DB
}

func NewRoutes(
	r *chi.Mux,
	dbConn *sql.DB,
) {

	route := &Route{
		mux:    r,
		dbConn: dbConn,
	}
	route.mux.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			route.getAuthRoutes(r)
			// route.getRoute(r)
			// route.getSwaggerRoutes(r)
			// route.getHealthRoutes(r)
		})

		r.Group(func(r chi.Router) {
			authMiddleware.NewMiddleware(r).AuthMiddleware()
			slog.Info("Auth OK")
			// route.getUserRoutes(r)
			// route.getRoleRoutes(r)
			// route.getPermissionRoutes(r)
		})
	})
}

func (r *Route) getAuthRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(r.dbConn)
	service := authService.NewAuthService(repo)
	handler := authHandler.NewAuthHandler(service)
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
	})
}

func (r *Route) getUserRoutes(router chi.Router) {
	repo := userRepository.NewUserRepository(r.dbConn)
	service := userService.NewUserService(repo)
	handler := userHandler.NewUserHandler(service)
	router.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
		r.Get("/{id}", handler.GetUser)
		r.Post("/", handler.CreateUser)
		r.Put("/{id}", handler.UpdateUser)
		r.Delete("/{id}", handler.DeleteUser)
	})
}

func (r *Route) getRoleRoutes(router chi.Router) {
	repo := roleRepository.NewRoleRepository(r.dbConn)
	service := roleService.NewRoleService(repo)
	handler := roleHandler.NewRoleHandler(service)

	repo2 := rolePermissionRepository.NewRolePermissionRepository(r.dbConn)
	service2 := rolePermissionService.NewRolePermissionService(repo2)
	handler2 := rolePermissionHandler.NewRolePermissionHandler(service2)

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

func (r *Route) getPermissionRoutes(router chi.Router) {
	repo := permissionRepository.NewPermissionRepository(r.dbConn)
	service := permissionService.NewPermissionService(repo)
	handler := permissionHandler.NewPermissionHandler(service)

	router.Route("/permission", func(r chi.Router) {
		r.Get("/", handler.GetPermissions)
		r.Get("/{id}", handler.GetPermission)
		r.Post("/", handler.CreatePermission)
		r.Put("/{id}", handler.UpdatePermission)
		r.Delete("/{id}", handler.DeletePermission)

	})
}

func (r *Route) getSwaggerRoutes(router chi.Router) {
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.NewEnv().GetPort()+"/api/v1/swagger/doc.json"),
	))
}

func (r *Route) getRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
	})
}

func (r *Route) getHealthRoutes(router chi.Router) {
	router.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
