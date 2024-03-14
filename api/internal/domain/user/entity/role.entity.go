package entity

import (
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)



type Role struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewRole(name string, internalName string, description string) (role *Role, err error) {
	role = &Role{
		Name:         name,
		InternalName: internalName,
		Description:  description,
	}
	notify := role.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}
	return
}

func (r *Role) Validate() (notify notification.ErrorsInterface) {

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

func (r *Role) SetInternalName(internalName string) {
	r.InternalName = internalName
}

func (r *Role) SetDescription(description string) {
	r.Description = description
}

func (r *Role) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

func (r *Role) SetUpdatedAt(updatedAt time.Time) {
	r.UpdatedAt = updatedAt
}
