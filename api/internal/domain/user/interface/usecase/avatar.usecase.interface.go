package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetAvatarUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error)
}

type GetAvatarsUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error)
}

type CreateAvatarUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error)
}

type UpdateAvatarUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error)
}

type DeleteAvatarUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.DeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error)
}
