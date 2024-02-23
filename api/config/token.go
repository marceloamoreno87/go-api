package config

import (
	"github.com/go-chi/jwtauth/v5"
)

type TokenAuth struct {
	Auth *jwtauth.JWTAuth
}

func NewTokenAuth() *TokenAuth {
	return &TokenAuth{
		Auth: jwtauth.New("HS256", []byte(NewEnv().GetJWTSecretKey()), nil),
	}
}

func (t *TokenAuth) GetAuth() *jwtauth.JWTAuth {
	return t.Auth
}

func (t *TokenAuth) GetJWTExpiresIn() string {
	return NewEnv().GetJWTExpiresIn()
}
