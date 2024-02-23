package entity

import (
	"time"
)

type AvatarInterface interface {
	GetID() int32
	GetSVG() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(int32)
	SetSVG(string)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}
