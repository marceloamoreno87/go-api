package main

import (
	"github.com/marceloamoreno/goapi/internal/infra/webserver"
)

// @title GO API
// @description This is a sample server for GO tools.
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
