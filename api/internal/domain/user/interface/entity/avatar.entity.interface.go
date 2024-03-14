package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
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
	Validate() notification.ErrorsInterface
}
