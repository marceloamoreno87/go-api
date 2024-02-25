package middleware

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/marceloamoreno/goapi/config"
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
	m.router.Use(jwtauth.Verifier(config.NewToken().GetAuth()))
	m.router.Use(jwtauth.Authenticator(config.NewToken().GetAuth()))
	slog.Info("Auth OK")
}
