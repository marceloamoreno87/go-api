package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/marceloamoreno/goapi/config"
)

type Auth struct {
	Token string `json:"token"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) NewToken(tokenAuth *jwtauth.JWTAuth, jwtExpiresIn string, id int64) (token string, err error) {
	jwtExpiresInInt, err := strconv.Atoi(jwtExpiresIn)
	if err != nil {
		return "", errors.New("Not Authorized")
	}

	_, tokenString, err := config.TokenAuth.Encode(map[string]interface{}{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresInInt)).Unix(),
	})

	if err != nil {
		return "", errors.New("Not Authorized")
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(tokenAuth *jwtauth.JWTAuth, token string) (id int64, err error) {
	tokenString, err := config.TokenAuth.Decode(token)

	idStr, bool := tokenString.Get("id")
	if !bool {
		return 0, errors.New("Not Authorized")
	}

	idInt64, err := strconv.ParseInt(fmt.Sprintf("%v", idStr), 10, 64)
	if err != nil {
		return 0, errors.New("Not Authorized")
	}

	return idInt64, nil
}

func (a *Auth) GetToken() string {
	return a.Token
}

func (a *Auth) SetToken(token string) {
	a.Token = token
}
