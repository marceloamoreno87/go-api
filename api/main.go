package main

import (
	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/internal/infra/webserver"
)

func init() {
	env := config.NewEnv()
	env.LoadEnv()

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
