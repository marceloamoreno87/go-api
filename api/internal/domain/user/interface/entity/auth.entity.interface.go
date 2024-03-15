package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type AuthInterface interface {
	SetID(id int32)
	SetUserID(userID int32)
	SetToken(token string)
	SetRefreshToken(refreshToken string)
	SetActive(active bool)
	SetExpiresIn(expiresIn int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	GetID() int32
	GetUserID() int32
	GetToken() string
	GetRefreshToken() string
	GetActive() bool
	GetExpiresIn() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetUser() UserInterface
	SetUser(user UserInterface)
	Validate() notification.ErrorsInterface
}
