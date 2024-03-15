package entityInterface

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type AuthInterface interface {
	SetID(id int32)
	SetUserID(userID int32)
	SetToken(token string)
	SetRefreshToken(refreshToken string)
	SetActive(active bool)
	SetTokenExpiresIn(tokenExpiresIn int32)
	SetRefreshTokenExpiresIn(refreshTokenExpiresIn int32)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	GetID() int32
	GetUserID() int32
	GetToken() string
	GetRefreshToken() string
	GetActive() bool
	GetTokenExpiresIn() int32
	GetRefreshTokenExpiresIn() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Validate() notification.ErrorsInterface
	IsValidToken() bool
	IsValidRefreshToken() bool
	GenerateToken()
}
