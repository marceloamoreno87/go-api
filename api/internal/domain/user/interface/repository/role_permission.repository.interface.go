package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type RolePermissionRepositoryInterface interface {
	GetRolePermissionsByRole(ctx context.Context, id int32) (output []entityInterface.RolePermissionInterface, err error)
	CreateRolePermission(ctx context.Context, rolePermission entityInterface.RolePermissionInterface) (err error)
	DeleteRolePermissionByRoleID(ctx context.Context, id int32) (err error)
}
