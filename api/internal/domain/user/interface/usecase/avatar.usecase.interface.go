package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetAvatarUseCaseInterface interface {
	Execute(input usecase.GetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error)
}

type GetAvatarsUseCaseInterface interface {
	Execute(input usecase.GetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
}

type CreateAvatarUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error)
}

type UpdateAvatarUseCaseInterface interface {
	Execute(input usecase.UpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error)
}

type DeleteAvatarUseCaseInterface interface {
	Execute(input usecase.DeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error)
}
