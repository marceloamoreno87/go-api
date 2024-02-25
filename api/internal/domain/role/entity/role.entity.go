package entity

import (
	"errors"
	"time"
)

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
	valid := role.Validate()
	if valid != nil {
		return nil, valid
	}

	return
}

func (r *Role) Validate() (err error) {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.InternalName == "" {
		return errors.New("internal Name is required")
	}
	if r.Description == "" {
		return errors.New("description is required")
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
