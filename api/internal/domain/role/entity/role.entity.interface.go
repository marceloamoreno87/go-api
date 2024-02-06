package entity

import (
	"time"
)

type RoleInterface interface {
	GetID() int32
	GetName() string
	GetInternalName() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Validate() error
	SetID(id int32)
	SetName(name string)
	SetInternalName(internal_name string)
	SetDescription(description string)
	SetCreatedAt(created_at time.Time)
	SetUpdatedAt(updated_at time.Time)
}
