package entity

import (
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type Auth struct {
	ID                    int32     `json:"id"`
	UserID                int32     `json:"user_id"`
	Token                 string    `json:"token"`
	RefreshToken          string    `json:"refresh_token"`
	Active                bool      `json:"active"`
	TokenExpiresIn        int32     `json:"token_expires_in"`
	RefreshTokenExpiresIn int32     `json:"refresh_token_expires_in"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func NewAuth(userId int32) (auth *Auth, err error) {
	auth = &Auth{
		UserID:                userId,
		Active:                true,
		TokenExpiresIn:        config.Jwt.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: config.Jwt.GetRefreshTokenExpiresIn(),
	}

	auth.GenerateToken()

	notify := auth.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (a *Auth) Validate() (notify notification.ErrorsInterface) {

	notify = notification.New()

	if a.UserID == 0 {
		notify.AddError("User is required", "auth.entity.user_id")
	}
	if a.Token == "" {
		notify.AddError("Token is required", "auth.entity.token")
	}
	if a.RefreshToken == "" {
		notify.AddError("RefreshToken is required", "auth.entity.refresh_token")
	}

	return
}

func (a *Auth) IsValidToken() (valid bool, err error) {
	valid = config.Jwt.Validate(a.Token)
	if !valid {
		err = errors.New("invalid token")
	}
	return
}

func (a *Auth) IsValidRefreshToken() (valid bool, err error) {
	valid = config.Jwt.Validate(a.RefreshToken)
	if !valid {
		err = errors.New("invalid refresh token")
	}
	return
}

func (a *Auth) GenerateToken() {
	claims := map[string]interface{}{
		"id": a.GetUserID(),
	}
	a.Token = config.Jwt.Generate(claims).GetToken()
	a.RefreshToken = config.Jwt.GenerateRefresh(map[string]interface{}{}).GetRefreshToken()
}

func (a *Auth) SetID(id int32) {
	a.ID = id
}

func (a *Auth) SetUserID(userID int32) {
	a.UserID = userID
}

func (a *Auth) SetToken(token string) {
	a.Token = token
}

func (a *Auth) SetRefreshToken(refreshToken string) {
	a.RefreshToken = refreshToken
}

func (a *Auth) SetActive(active bool) {
	a.Active = active
}

func (a *Auth) SetTokenExpiresIn(tokenExpiresIn int32) {
	a.TokenExpiresIn = tokenExpiresIn
}

func (a *Auth) SetRefreshTokenExpiresIn(refreshTokenExpiresIn int32) {
	a.RefreshTokenExpiresIn = refreshTokenExpiresIn
}

func (a *Auth) SetCreatedAt(createdAt time.Time) {
	a.CreatedAt = createdAt
}

func (a *Auth) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

func (a *Auth) GetID() int32 {
	return a.ID
}

func (a *Auth) GetUserID() int32 {
	return a.UserID
}

func (a *Auth) GetToken() string {
	return a.Token
}

func (a *Auth) GetRefreshToken() string {
	return a.RefreshToken
}

func (a *Auth) GetActive() bool {
	return a.Active
}

func (a *Auth) GetTokenExpiresIn() int32 {
	return a.TokenExpiresIn
}

func (a *Auth) GetRefreshTokenExpiresIn() int32 {
	return a.RefreshTokenExpiresIn
}

func (a *Auth) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *Auth) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}
