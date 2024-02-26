package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CorsMiddleware struct {
	router chi.Router
}

func NewCorsMiddleware(router chi.Router) *CorsMiddleware {
	return &CorsMiddleware{
		router: router,
	}
}

func (m *CorsMiddleware) CorsMiddleware() {
	m.router.Use(middleware.AllowContentType("application/json"))
	m.router.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	m.router.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE"))
	m.router.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
}
