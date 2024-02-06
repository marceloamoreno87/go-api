package entity

import (
	"time"
)

type PermissionInterface interface {
	GetID() int
	GetName() string
	GetInternalName() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id int)
	SetName(name string)
	SetInternalName(internalName string)
	SetDescription(description string)
}
