package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type PermissionInterface interface {
	GetID() int32
	GetName() string
	GetInternalName() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id int32)
	SetName(name string)
	SetInternalName(internalName string)
	SetDescription(description string)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	Validate() (notify notification.ErrorsInterface)
}