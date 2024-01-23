package config

import "github.com/go-chi/jwtauth/v5"

var TokenAuth *jwtauth.JWTAuth

func NewTokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(Environment.JWTSecretKey), nil)
}
