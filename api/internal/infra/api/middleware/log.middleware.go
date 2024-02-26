package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type LogMiddleware struct {
	router chi.Router
}

func NewLogMiddleware(router chi.Router) *LogMiddleware {
	return &LogMiddleware{
		router: router,
	}
}

func (m *LogMiddleware) LogMiddleware() {
	m.router.Use(middleware.Logger)
}
