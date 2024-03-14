package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type AuthMiddleware struct {
	router chi.Router
}

func NewMiddleware(router chi.Router) *AuthMiddleware {
	return &AuthMiddleware{
		router: router,
	}
}

func (m *AuthMiddleware) AuthMiddleware(jwt *jwtauth.JWTAuth) {
	m.router.Use(jwtauth.Verifier(jwt))
	m.router.Use(jwtauth.Authenticator(jwt))
}
