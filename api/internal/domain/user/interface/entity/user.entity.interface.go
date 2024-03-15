package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type UserInterface interface {
	GetID() int32
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRoleID() int32
	GetAvatarID() int32
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id int32)
	SetName(name string)
	SetEmail(email string)
	SetPassword(password string)
	SetActive(active bool)
	SetRoleID(roleID int32)
	SetAvatarID(avatarID int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	Validate() (notify notification.ErrorsInterface)
	ComparePassword(password string) (b bool)
	HashPassword()
}
