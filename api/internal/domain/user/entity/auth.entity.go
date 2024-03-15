package entity

import (
	"errors"
	"time"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type Auth struct {
	ID           int32                         `json:"id"`
	UserID       int32                         `json:"user_id"`
	Token        string                        `json:"token"`
	RefreshToken string                        `json:"refresh_token"`
	Active       bool                          `json:"active"`
	ExpiresIn    int32                         `json:"expires_in"`
	CreatedAt    time.Time                     `json:"created_at"`
	UpdatedAt    time.Time                     `json:"updated_at"`
	User         entityInterface.UserInterface `json:"user"`
}

func NewAuth(userId int32, token string, refreshToken string) (auth entityInterface.AuthInterface, err error) {
	newAuth := &Auth{
		UserID:       userId,
		Token:        token,
		RefreshToken: refreshToken,
	}

	notify := newAuth.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (u *Auth) Validate() (notify notification.ErrorsInterface) {

	notify = notification.New()

	if u.UserID == 0 {
		notify.AddError("User is required", "auth.entity.user_id")
	}
	if u.Token == "" {
		notify.AddError("Token is required", "auth.entity.token")
	}
	if u.RefreshToken == "" {
		notify.AddError("RefreshToken is required", "auth.entity.refresh_token")
	}

	return
}

func (a *Auth) SetID(id int32) {
	a.ID = id
}

func (a *Auth) SetUserID(userID int32) {
	a.UserID = userID
}

func (a *Auth) SetUser(user entityInterface.UserInterface) {
	a.User = user
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

func (a *Auth) SetExpiresIn(expiresIn int32) {
	a.ExpiresIn = expiresIn
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

func (a *Auth) GetUser() entityInterface.UserInterface {
	return a.User
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

func (a *Auth) GetExpiresIn() int32 {
	return a.ExpiresIn
}

func (a *Auth) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *Auth) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}
