package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type RoleRepositoryInterface interface {
	CreateRole(ctx context.Context, role entityInterface.RoleInterface) (err error)
	GetRole(ctx context.Context, id int32) (output entityInterface.RoleInterface, err error)
	GetRoleByInternalName(ctx context.Context, internalName string) (output entityInterface.RoleInterface, err error)
	GetRoles(ctx context.Context, limit int32, offset int32) (output []entityInterface.RoleInterface, err error)
	UpdateRole(ctx context.Context, role entityInterface.RoleInterface, id int32) (err error)
	DeleteRole(ctx context.Context, id int32) (err error)
}
