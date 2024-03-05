package entity

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"time"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type UserValidationInterface interface {
	Validate() (notify *notification.Errors)
}

type UserValidation struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Hash      string    `json:"hash"`
	ExpiresIn string    `json:"expires_in"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserValidation(user *User) (userValidation *UserValidation, err error) {
	userValidation = &UserValidation{
		User:      user,
		UserID:    user.GetID(),
		ExpiresIn: config.Environment.GetValidationExpiresIn(),
	}
	return
}

func (u *UserValidation) Validate() (notify *notification.Errors) {
	notify = notification.New()
	if u.UserID == 0 {
		notify.AddError("User is required", "user_validation.entity.user")
	}
	return
}

func (u *UserValidation) GenerateHash() (err error) {
	userJson, err := json.Marshal(u.User)
	if err != nil {
		return
	}
	mergeStr := string(userJson) + config.Environment.GetValidationSecretKey()
	sha512.New().Write([]byte(mergeStr))
	u.Hash = fmt.Sprintf("%x", sha512.Sum512([]byte(mergeStr)))
	return
}

func (u *UserValidation) GetID() int32 {
	return u.ID
}

func (u *UserValidation) GetUserID() int32 {
	return u.UserID
}

func (u *UserValidation) GetHash() string {
	return u.Hash
}

func (u *UserValidation) GetExpiresIn() string {
	return u.ExpiresIn
}

func (u *UserValidation) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserValidation) GetUser() *User {
	return u.User
}

func (u *UserValidation) SetID(id int32) {
	u.ID = id
}

func (u *UserValidation) SetUserID(userID int32) {
	u.UserID = userID
}

func (u *UserValidation) SetHash(hash string) {
	u.Hash = hash
}

func (u *UserValidation) SetExpiresIn(expiresIn string) {
	u.ExpiresIn = expiresIn
}
