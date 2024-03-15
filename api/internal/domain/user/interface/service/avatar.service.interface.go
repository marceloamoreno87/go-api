package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(limit int32, offset int32) (output []usecase.GetAvatarsOutputDTO, err error)
	CreateAvatar(body io.ReadCloser) (output usecase.CreateAvatarOutputDTO, err error)
	UpdateAvatar(id int32, body io.ReadCloser) (output usecase.UpdateAvatarOutputDTO, err error)
	DeleteAvatar(id int32) (output usecase.DeleteAvatarOutputDTO, err error)
}
