package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetRolePermissionsUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetRolePermissionsInputDTO) (output []usecase.GetRolePermissionsOutputDTO, err error)
}

type CreateRolePermissionUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
}

type DeleteRolePermissionByRoleIDUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.DeleteRolePermissionByRoleIDInputDTO) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
}
