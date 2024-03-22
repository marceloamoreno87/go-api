package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetPermissionUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetPermissionInputDTO) (output usecase.GetPermissionOutputDTO, err error)
}

type GetPermissionsUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error)
}

type CreatePermissionUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.CreatePermissionInputDTO) (output usecase.CreatePermissionOutputDTO, err error)
}

type UpdatePermissionUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.UpdatePermissionInputDTO) (output usecase.UpdatePermissionOutputDTO, err error)
}

type DeletePermissionUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.DeletePermissionInputDTO) (output usecase.DeletePermissionOutputDTO, err error)
}

type GetPermissionByInternalNameUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetPermissionByInternalNameInputDTO) (output usecase.GetPermissionByInternalNameOutputDTO, err error)
}
