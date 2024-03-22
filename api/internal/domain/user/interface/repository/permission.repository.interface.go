package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type PermissionRepositoryInterface interface {
	CreatePermission(ctx context.Context, permission entityInterface.PermissionInterface) (err error)
	GetPermission(ctx context.Context, id int32) (output entityInterface.PermissionInterface, err error)
	GetPermissions(ctx context.Context, limit int32, offset int32) (output []entityInterface.PermissionInterface, err error)
	UpdatePermission(ctx context.Context, permission entityInterface.PermissionInterface, id int32) (err error)
	DeletePermission(ctx context.Context, id int32) (err error)
	GetPermissionByInternalName(ctx context.Context, internalName string) (output entityInterface.PermissionInterface, err error)
}
