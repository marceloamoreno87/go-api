package main

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/infra/webserver"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

func init() {
	config.NewEnv()
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
	m := mail.NewMail()
	m.SetFrom("marceloamoreno87@gmail.com")
	m.SetTo([]string{"marceloamoreno87@gmail.com"})
	m.SetSubject("teste")
	m.SetBody("teste")
	err := m.Send()
	if err != nil {
		panic(err)
	}
	webserver.Bootstrap()
}
