package entity

import (
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type UserValidation struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Hash      string    `json:"hash"`
	ExpiresIn int32     `json:"expires_in"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserValidation(userId int32) (userValidation entityInterface.UserValidationInterface, err error) {
	userValidation = &UserValidation{
		UserID:    userId,
		Used:      false,
		ExpiresIn: int32(time.Now().Add(time.Second * time.Duration(helper.StrToInt32(config.Environment.GetValidationExpiresIn()))).Unix()),
	}

	err = userValidation.GenerateHash()
	if err != nil {
		return nil, err
	}

	notify := userValidation.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (u *UserValidation) Validate() (notify notification.ErrorsInterface) {
	notify = notification.New()

	if u.UserID == 0 {
		notify.AddError("UserID", "UserID is required")
	}

	if u.Hash == "" {
		notify.AddError("Hash", "Hash is required")
	}

	return
}

func (u *UserValidation) GenerateHash() (err error) {
	userJson, err := json.Marshal(u)
	if err != nil {
		return
	}
	mergeStr := string(userJson) + config.Environment.GetValidationSecretKey()
	sha512.New().Write([]byte(mergeStr))
	u.Hash = fmt.Sprintf("%x", sha512.Sum512([]byte(mergeStr)))
	return
}

func (u *UserValidation) ValidateHashExpiresIn() bool {
	return u.ExpiresIn > int32(time.Now().Unix())
}

func (u *UserValidation) GetUsed() bool {
	return u.Used
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

func (u *UserValidation) GetExpiresIn() int32 {
	return u.ExpiresIn
}

func (u *UserValidation) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserValidation) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *UserValidation) SetID(id int32) {
	u.ID = id
}

func (u *UserValidation) SetUsed(used bool) {
	u.Used = used
}

func (u *UserValidation) SetUserID(userID int32) {
	u.UserID = userID
}

func (u *UserValidation) SetHash(hash string) {
	u.Hash = hash
}

func (u *UserValidation) SetExpiresIn(expiresIn int32) {
	u.ExpiresIn = expiresIn
}
