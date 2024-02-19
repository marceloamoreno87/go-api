package entity

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
)

type UserInterface interface {
	GetID() int32
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRoleID() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetRole() *entity.Role
	SetID(id int32)
	SetName(name string)
	SetEmail(email string)
	SetPassword(password string)
	SetRoleID(roleId int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	SetRole(role *entity.Role)
	Validate() (err error)
	ComparePassword(password string) (b bool)
}
