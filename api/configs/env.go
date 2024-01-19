package configs

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
)

var Environment *Env
var TokenAuth *jwtauth.JWTAuth

type Env struct {
	NameProject   string
	Host          string
	Port          string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecretKey  string
	JWTExperiesIn string
}

func NewEnv() *Env {
	return &Env{
		NameProject:   os.Getenv("NAME_PROJECT"),
		Host:          os.Getenv("HOST"),
		Port:          os.Getenv("PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		JWTSecretKey:  os.Getenv("JWT_SECRET_KEY"),
		JWTExperiesIn: os.Getenv("JWT_EXPERIES_IN"),
	}
}

func NewTokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(Environment.JWTSecretKey), nil)
}

func (env *Env) LoadEnv() {
	Environment = NewEnv()
	TokenAuth = NewTokenAuth()
}
