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
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetRole() RoleInterface
	GetAvatar() AvatarInterface
	SetID(id int32)
	SetName(name string)
	SetEmail(email string)
	SetPassword(password string)
	SetRoleID(roleID int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	SetRole(role RoleInterface)
	SetAvatarID(avatarID int32)
	SetAvatar(avatar AvatarInterface)
	Validate() (notify notification.ErrorsInterface)
	ComparePassword(password string) (b bool)
	HashPassword()
	SetActive(active bool)
	GetActive() bool
	GenerateToken()
	GetToken() string
}
