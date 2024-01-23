package config

import (
	"os"
)

var Environment *Env

type Env struct {
	NameProject  string
	Host         string
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecretKey string
	JWTExpiresIn string
}

func NewEnv() *Env {
	return &Env{
		NameProject:  os.Getenv("NAME_PROJECT"),
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		JWTExpiresIn: os.Getenv("JWT_EXPIRES_IN"),
	}
}

func (env *Env) LoadEnv() {
	Environment = NewEnv()
}
