package webserver

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/infra/api/routes"
)

type Server struct {
	mux    config.MuxInterface
	jwt    config.JWTAuthInterface
}

func NewServer() *Server {
	return &Server{
		mux:    config.M,
		jwt:    config.Jwt,
	}
}

func (s *Server) Start() {
	routes.NewRoutes(s.mux, s.jwt)
	port := config.Environment.GetPort()
	slog.Info("Server started on port http://localhost:" + port + "/api/v1")
	slog.Info("Swagger started on port http://localhost:" + port + "/api/v1/swagger/index.html")
	slog.Info("Health started on port http://localhost:" + port + "/api/v1/health")
	if err := http.ListenAndServe(":"+port, s.mux.GetMux()); err != nil {
		panic(err)
	}
}
