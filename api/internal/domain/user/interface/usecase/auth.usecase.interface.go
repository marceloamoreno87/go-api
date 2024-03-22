package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type CreateAuthUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateAuthInputDTO) (output usecase.CreateAuthOutputDTO, err error)
}
type LoginUserUseCaseInterface interface {
	Execute(input usecase.LoginUserInputDTO) (output usecase.LoginUserOutputDTO, err error)
}

type GetAuthByUserIDUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetAuthByUserIDInputDTO) (output usecase.GetAuthByUserIDOutputDTO, err error)
}

type GetAuthByTokenUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetAuthByTokenInputDTO) (output usecase.GetAuthByTokenOutputDTO, err error)
}

type UpdateAuthRevokeUseCaseInterface interface {
	Execute(ictx context.Context, nput usecase.UpdateAuthRevokeInputDTO) (output usecase.UpdateAuthRevokeOutputDTO, err error)
}

type GetAuthByRefreshTokenUseCase interface {
	Execute(ctx context.Context, input usecase.GetAuthByRefreshTokenInputDTO) (output usecase.GetAuthByRefreshTokenOutputDTO, err error)
}

type UpdateUserActiveUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateUserActiveInputDTO) (err error)
}

type UpdateUserValidationUsedUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateUserValidationUsedInputDTO) (err error)
}
