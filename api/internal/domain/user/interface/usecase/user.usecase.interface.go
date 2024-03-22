package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetUserByEmailUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetUserByEmailInputDTO) (output usecase.GetUserByEmailOutputDTO, err error)
}

type CreateUserUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateUserInputDTO) (output usecase.CreateUserOutputDTO, err error)
}

type GetUserUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetUserInputDTO) (output usecase.GetUserOutputDTO, err error)
}

type GetUsersUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetUsersInputDTO) (output []usecase.GetUsersOutputDTO, err error)
}

type UpdateUserUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateUserInputDTO) (output usecase.UpdateUserOutputDTO, err error)
}

type DeleteUserUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.DeleteUserInputDTO) (output usecase.DeleteUserOutputDTO, err error)
}

type UpdateUserPasswordUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateUserPasswordInputDTO) (output usecase.UpdateUserPasswordOutputDTO, err error)
}

type CreateUserValidationUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateUserValidationInputDTO) (output usecase.CreateUserValidationOutputDTO, err error)
}
