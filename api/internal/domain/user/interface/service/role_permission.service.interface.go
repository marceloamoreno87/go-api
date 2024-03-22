package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(ctx context.Context, input request.RequestGetRolePermission) (output []usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(ctx context.Context, input request.RequestCreateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error)
	DeleteRolePermissionByRoleID(ctx context.Context, input request.RequestDeleteRolePermissionByRoleID) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
	UpdateRolePermission(ctx context.Context, input request.RequestUpdateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error)
}
