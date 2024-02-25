package entity

import (
	"strconv"
	"time"

	"github.com/marceloamoreno/goapi/config"
)

type Auth struct {
	Token string `json:"token"`
	ID    int32  `json:"id"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) NewToken(token config.TokenInterface, jwtExpiresIn string, id int32) (err error) {
	jwtExpiresInInt, err := strconv.Atoi(jwtExpiresIn)
	if err != nil {
		return
	}

	_, tokenString, err := token.GetAuth().Encode(map[string]interface{}{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresInInt)).Unix(),
	})

	if err != nil {
		return
	}

	a.SetToken(tokenString)
	return
}

func (a *Auth) RefreshToken(token config.TokenInterface, lastToken string) (err error) {
	tokenString, err := token.GetAuth().Decode(lastToken)
	if err != nil {
		return
	}

	idStr, bool := tokenString.Get("id")
	if !bool {
		return
	}

	idInt32 := int32(idStr.(float64))
	a.SetID(idInt32)
	return
}

func (a *Auth) GetToken() string {
	return a.Token
}

func (a *Auth) SetToken(token string) {
	a.Token = token
}

func (a *Auth) GetID() int32 {
	return a.ID
}

func (a *Auth) SetID(id int32) {
	a.ID = id
}
