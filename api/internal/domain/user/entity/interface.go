package entity

import (
	"time"
)

type UserInterface interface {
	Validate() error
	IsEmailValid() (bool, error)
	ComparePassword(password string) bool
	GetID() int64
	GetName() string
	GetEmail() string
	GetPassword() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id int64)
	SetName(name string)
	SetEmail(email string)
	SetPassword(password string)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
}
