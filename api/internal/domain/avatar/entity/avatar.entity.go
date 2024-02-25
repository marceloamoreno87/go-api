package entity

import (
	"errors"
	"time"
)

type Avatar struct {
	ID        int32     `json:"id"`
	SVG       string    `json:"svg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAvatar(SVG string) (avatar *Avatar, err error) {
	avatar = &Avatar{
		SVG: SVG,
	}
	valid := avatar.Validate()
	if valid != nil {
		return nil, valid
	}

	return
}

func (u *Avatar) Validate() (err error) {
	if u.SVG == "" {
		return errors.New("SVG is required")
	}
	return
}

func (u *Avatar) GetID() int32 {
	return u.ID
}

func (u *Avatar) GetSVG() string {
	return u.SVG
}

func (u *Avatar) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *Avatar) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *Avatar) SetID(id int32) {
	u.ID = id
}

func (u *Avatar) SetSVG(SVG string) {
	u.SVG = SVG
}

func (u *Avatar) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *Avatar) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
