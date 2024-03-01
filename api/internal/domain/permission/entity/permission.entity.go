package entity

import (
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type Permission struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewPermission(name string, internalName string, description string) (permission *Permission, err error) {
	permission = &Permission{
		Name:         name,
		InternalName: internalName,
		Description:  description,
	}

	notify := permission.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (p *Permission) Validate() (notify *notification.Errors) {

	notify = notification.New()

	if p.Name == "" {
		notify.AddError("Name is required", "permission.entity.name")
	}
	if p.InternalName == "" {
		notify.AddError("Internal name is required", "permission.entity.internal_name")
	}
	if p.Description == "" {
		notify.AddError("Description is required", "permission.entity.description")
	}
	return
}

func (p *Permission) GetID() int32 {
	return p.ID
}

func (p *Permission) GetName() string {
	return p.Name
}

func (p *Permission) GetInternalName() string {
	return p.InternalName
}

func (p *Permission) GetDescription() string {
	return p.Description
}

func (p *Permission) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Permission) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *Permission) SetID(id int32) {
	p.ID = id
}

func (p *Permission) SetName(name string) {
	p.Name = name
}

func (p *Permission) SetInternalName(internalName string) {
	p.InternalName = internalName
}

func (p *Permission) SetDescription(description string) {
	p.Description = description
}

func (p *Permission) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Permission) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}
