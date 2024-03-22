package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetRoleUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetRoleInputDTO) (output usecase.GetRoleOutputDTO, err error)
}

type GetRolesUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetRolesInputDTO) (output []usecase.GetRolesOutputDTO, err error)
}

type CreateRoleUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateRoleInputDTO) (output usecase.CreateRoleOutputDTO, err error)
}

type UpdateRoleUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdateRoleInputDTO) (output usecase.UpdateRoleOutputDTO, err error)
}

type DeleteRoleUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.DeleteRoleInputDTO) (output usecase.DeleteRoleOutputDTO, err error)
}

type NewGetRoleByInternalNameUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetRoleByInternalNameInputDTO) (output usecase.GetRoleByInternalNameOutputDTO, err error)
}
