package main

import (
	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/internal/infra/webserver"
)

// @title GO API
// @description This is a sample server celler server.
// @version v1
// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.api {
// 	"Authorization": {
// 		"type": "Bearer",
// 		"name": "Authorization",
// 		"in": "header"
// 	}

func init() {
	env := config.NewEnv()
	env.LoadEnv()

}

func main() {
	webserver.StartServer()
}
