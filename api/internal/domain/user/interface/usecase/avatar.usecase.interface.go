package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreateAvatarUseCaseInterface interface {
	Execute(input usecase.CreateAvatarInputDTO) (err error)
}

type GetAvatarUseCaseInterface interface {
	Execute(id int32) (output usecase.GetAvatarOutputDTO, err error)
}

type GetAvatarsUseCaseInterface interface {
	Execute(input usecase.GetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
}

type UpdateAvatarUseCaseInterface interface {
	Execute(id int32, input usecase.UpdateAvatarInputDTO) (err error)
}

type DeleteAvatarUseCaseInterface interface {
	Execute(id int32) (err error)
}
