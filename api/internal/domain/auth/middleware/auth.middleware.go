package middleware

import (
	"github.com/go-chi/chi/v5"
)

type AuthMiddleware struct {
	router chi.Router
}

func NewMiddleware(router chi.Router) *AuthMiddleware {
	return &AuthMiddleware{
		router: router,
	}
}

func (m *AuthMiddleware) AuthMiddleware() {
}
