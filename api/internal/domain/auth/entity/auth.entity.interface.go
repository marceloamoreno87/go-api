package entity

import "github.com/go-chi/jwtauth/v5"

type AuthInterface interface {
	NewToken(tokenAuth *jwtauth.JWTAuth, jwtExpiresIn string, id int64) error
	RefreshToken(tokenAuth *jwtauth.JWTAuth, token string) error
	GetToken() string
	SetToken(token string)
	GetId() int64
	SetId(id int64)
}
