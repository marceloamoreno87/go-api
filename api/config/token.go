package config

import (
	"github.com/go-chi/jwtauth/v5"
)

type TokenInterface interface {
	GetAuth() *jwtauth.JWTAuth
	GetJWTExpiresIn() string
}

type Token struct {
	auth *jwtauth.JWTAuth
}

func NewToken() *Token {
	return &Token{
		auth: jwtauth.New("HS256", []byte(NewEnv().GetJWTSecretKey()), nil),
	}
}

func (t *Token) GetAuth() *jwtauth.JWTAuth {
	return t.auth
}

func (t *Token) GetJWTExpiresIn() string {
	return NewEnv().GetJWTExpiresIn()
}
