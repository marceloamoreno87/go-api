package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionServiceInterface interface {
	GetPermission(ctx context.Context, input request.RequestGetPermission) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(ctx context.Context, input request.RequestGetPermissions) (output []usecase.GetPermissionsOutputDTO, err error)
	CreatePermission(ctx context.Context, input request.RequestCreatePermission) (output usecase.CreatePermissionOutputDTO, err error)
	UpdatePermission(ctx context.Context, input request.RequestUpdatePermission) (output usecase.UpdatePermissionOutputDTO, err error)
	DeletePermission(ctx context.Context, input request.RequestDeletePermission) (output usecase.DeletePermissionOutputDTO, err error)
}
