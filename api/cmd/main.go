package main

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/infra/webserver"
)

func init() {
	config.NewEnv()
	config.NewTokenAuth()
	slog.Info("Environment OK")
}

// @title GO API
// @description This is a sample server for GO API.
// @version v1
// @host localhost:3000
// @BasePath /api/v1
// @schemes http
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
func main() {
	webserver.StartServer()
}
