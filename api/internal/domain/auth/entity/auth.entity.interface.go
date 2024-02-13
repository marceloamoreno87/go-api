package entity

import "github.com/go-chi/jwtauth/v5"

type AuthInterface interface {
	NewToken(tokenAuth *jwtauth.JWTAuth, jwtExpiresIn string, id int32) (err error)
	RefreshToken(tokenAuth *jwtauth.JWTAuth, token string) (err error)
	GetToken() string
	SetToken(token string)
	GetId() int32
	SetId(id int32)
}
