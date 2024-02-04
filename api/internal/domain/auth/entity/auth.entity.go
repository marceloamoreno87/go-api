package entity

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type Auth struct {
	Token string `json:"token"`
	Id    int64  `json:"id"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) NewToken(tokenAuth *jwtauth.JWTAuth, jwtExpiresIn string, id int64) error {
	jwtExpiresInInt, err := strconv.Atoi(jwtExpiresIn)
	if err != nil {
		return errors.New("not authorized")
	}

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresInInt)).Unix(),
	})

	if err != nil {
		return errors.New("not authorized")
	}

	a.SetToken(tokenString)
	return nil
}

func (a *Auth) RefreshToken(tokenAuth *jwtauth.JWTAuth, token string) error {
	tokenString, err := tokenAuth.Decode(token)
	if err != nil {
		return errors.New("not authorized")
	}

	idStr, bool := tokenString.Get("id")
	if !bool {
		return errors.New("not authorized")
	}

	idInt64, err := strconv.ParseInt(fmt.Sprintf("%v", idStr), 10, 64)
	if err != nil {
		return errors.New("not authorized")
	}

	a.SetId(idInt64)
	return nil
}

func (a *Auth) GetToken() string {
	return a.Token
}

func (a *Auth) SetToken(token string) {
	a.Token = token
}

func (a *Auth) GetId() int64 {
	return a.Id
}

func (a *Auth) SetId(id int64) {
	a.Id = id
}
