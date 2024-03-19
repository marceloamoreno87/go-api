package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	GetAvatar(input request.RequestGetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(input request.RequestGetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
	CreateAvatar(input request.RequestCreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error)
	UpdateAvatar(input request.RequestUpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error)
	DeleteAvatar(input request.RequestDeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error)
}
