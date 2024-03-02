package jwt

import (
	"log/slog"

	"github.com/go-chi/jwtauth/v5"
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type JWTAuth struct {
	Token string
}

func New(claims map[string]interface{}) *JWTAuth {
	claims["exp"] = helper.StrToInt32(config.Environment.GetJWTExpiresIn())
	_, token, err := jwtauth.New("HS256", []byte(config.Environment.GetJWTSecretKey()), nil).Encode(claims)
	if err != nil {
		slog.Info("err", err)
	}

	return &JWTAuth{
		Token: token,
	}
}
