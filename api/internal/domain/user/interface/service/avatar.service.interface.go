package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	GetAvatar(input request.RequestGetAvatar) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(input request.RequestGetAvatars) (output []usecase.GetAvatarsOutputDTO, err error)
	CreateAvatar(input request.RequestCreateAvatar) (output usecase.CreateAvatarOutputDTO, err error)
	UpdateAvatar(input request.RequestUpdateAvatar) (output usecase.UpdateAvatarOutputDTO, err error)
	DeleteAvatar(input request.RequestDeleteAvatar) (output usecase.DeleteAvatarOutputDTO, err error)
}
