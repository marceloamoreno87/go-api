package entity

import (
	"errors"
	"time"
)

// Permission Entity
type Permission struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewPermission(name, internalName, description string) (permission *Permission, err error) {
	permission = &Permission{
		Name:         name,
		InternalName: internalName,
		Description:  description,
	}

	err = permission.Validate()
	if err != nil {
		return nil, err
	}

	return
}

func (p *Permission) Validate() (err error) {
	if p.Name == "" {
		return errors.New("Name is required")
	}
	if p.InternalName == "" {
		return errors.New("Internal Name is required")
	}
	if p.Description == "" {
		return errors.New("Description is required")
	}
	return
}

func (p *Permission) GetID() int {
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

func (p *Permission) SetID(id int) {
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
