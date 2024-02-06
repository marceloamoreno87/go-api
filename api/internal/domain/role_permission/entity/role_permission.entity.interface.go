package entity

import (
	"time"
)

type UserInterface interface {
	Validate() error
	IsEmailValid() (bool, error)
	ComparePassword(password string) bool
	GetID() int32
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRoleId() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id int32)
	SetName(name string)
	SetEmail(email string)
	SetPassword(password string)
	SetRoleId(roleId int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
}
