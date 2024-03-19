package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	GetAvatar(input service.RequestGetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(input service.RequestGetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
	CreateAvatar(input service.RequestCreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error)
	UpdateAvatar(input service.RequestUpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error)
	DeleteAvatar(input service.RequestDeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error)
}
