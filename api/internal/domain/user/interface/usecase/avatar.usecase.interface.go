package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type NewGetAvatarUseCaseInterface interface {
	Execute(input usecase.GetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error)
}

type NewGetAvatarsUseCaseInterface interface {
	Execute(input usecase.GetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
}

type NewCreateAvatarUseCaseInterface interface {
	Execute(input usecase.CreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error)
}

type NewUpdateAvatarUseCaseInterface interface {
	Execute(input usecase.UpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error)
}

type NewDeleteAvatarUseCaseInterface interface {
	Execute(input usecase.DeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error)
}
