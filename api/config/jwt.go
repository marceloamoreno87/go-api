package config

import (
	"log/slog"
	"time"

	"github.com/go-chi/jwtauth/v5"
	JWTChi "github.com/go-chi/jwtauth/v5"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type JWTAuthInterface interface {
	Generate(claims map[string]interface{}) *JWTAuth
	GenerateRefresh(claims map[string]interface{}) *JWTAuth
	GetToken() string
	GetJwtAuth() *jwtauth.JWTAuth
	GetTokenExpiresIn() int32
	GetRefreshToken() string
	GetRefreshTokenExpiresIn() int32
	Validate(token string) (b bool)
}

type JWTAuth struct {
	Token        string
	RefreshToken string
	JwtAuth      *jwtauth.JWTAuth
}

var Jwt JWTAuthInterface

func NewJWT() {
	jwt := &JWTAuth{
		JwtAuth: JWTChi.New("HS256", []byte(Environment.GetJWTSecretKey()), nil),
	}
	Jwt = jwt
}

func (j *JWTAuth) Generate(claims map[string]interface{}) *JWTAuth {
	claims["exp"] = j.GetTokenExpiresIn()
	_, token, err := j.JwtAuth.Encode(claims)
	if err != nil {
		slog.Info("err", err)
	}
	j.Token = token
	return j
}

func (j *JWTAuth) GenerateRefresh(claims map[string]interface{}) *JWTAuth {
	claims["exp"] = j.GetRefreshTokenExpiresIn()
	_, token, err := j.JwtAuth.Encode(claims)
	if err != nil {
		slog.Info("err", err)
	}
	j.RefreshToken = token
	return j
}

func (j *JWTAuth) Validate(token string) (b bool) {
	t, err := j.JwtAuth.Decode(token)
	if err != nil {
		return false
	}
	if t.Expiration().Unix() > time.Now().Unix() {
		return true
	}
	return false
}

func (j *JWTAuth) GetToken() string {
	return j.Token
}

func (j *JWTAuth) GetRefreshToken() string {
	return j.RefreshToken
}

func (j *JWTAuth) GetJwtAuth() *jwtauth.JWTAuth {
	return j.JwtAuth
}

func (j *JWTAuth) GetTokenExpiresIn() int32 {
	return int32(time.Now().Add(time.Second * time.Duration(helper.StrToInt32(Environment.GetJWTExpiresIn()))).Unix())
}

func (j *JWTAuth) GetRefreshTokenExpiresIn() int32 {
	return int32(time.Now().Add(time.Second * time.Duration(helper.StrToInt32(Environment.GetJWTExpiresIn())) * 2).Unix())
}
