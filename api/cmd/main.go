package main

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/infra/webserver"
)

func init() {
	config.NewEnv()
	config.NewDatabase()
	config.NewJWT()
	config.NewMux()
	config.NewSqlc(config.Db)
}

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
	webserver.NewServer().Start()
}
