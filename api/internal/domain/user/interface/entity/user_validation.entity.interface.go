package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type UserValidationInterface interface {
	Validate() (notify notification.ErrorsInterface)
	GenerateHash() (err error)
	GetID() int32
	GetUserID() int32
	GetHash() string
	GetExpiresIn() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetUsed() bool
	SetID(id int32)
	SetUserID(userID int32)
	SetHash(hash string)
	SetUsed(used bool)
	SetExpiresIn(expiresIn int32)
	ValidateHashExpiresIn() bool
}
