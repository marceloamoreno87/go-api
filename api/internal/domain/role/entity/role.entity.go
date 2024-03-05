package entity

import (
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
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

type Role struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewRole(name string, internal_name string, description string) (role *Role, err error) {
	role = &Role{
		Name:         name,
		InternalName: internal_name,
		Description:  description,
	}
	notify := role.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}
	return
}

func (r *Role) Validate() (notify *notification.Errors) {

	notify = notification.New()

	if r.Name == "" {
		notify.AddError("Name is required", "role.entity.name")
	}
	if r.InternalName == "" {
		notify.AddError("Internal name is required", "role.entity.internal_name")
	}
	if r.Description == "" {
		notify.AddError("Description is required", "role.entity.description")
	}
	return
}

func (r *Role) GetID() int32 {
	return r.ID
}

func (r *Role) GetName() string {
	return r.Name
}

func (r *Role) GetInternalName() string {
	return r.InternalName
}

func (r *Role) GetDescription() string {
	return r.Description
}

func (r *Role) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r *Role) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

func (r *Role) SetID(id int32) {
	r.ID = id
}

func (r *Role) SetName(name string) {
	r.Name = name
}

func (r *Role) SetInternalName(internal_name string) {
	r.InternalName = internal_name
}

func (r *Role) SetDescription(description string) {
	r.Description = description
}

func (r *Role) SetCreatedAt(created_at time.Time) {
	r.CreatedAt = created_at
}

func (r *Role) SetUpdatedAt(updated_at time.Time) {
	r.UpdatedAt = updated_at
}
