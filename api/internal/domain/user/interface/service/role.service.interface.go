package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	GetRole(ctx context.Context, input request.RequestGetRole) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(ctx context.Context, input request.RequestGetRoles) (output []usecase.GetRolesOutputDTO, err error)
	CreateRole(ctx context.Context, input request.RequestCreateRole) (output usecase.CreateRoleOutputDTO, err error)
	UpdateRole(ctx context.Context, input request.RequestUpdateRole) (output usecase.UpdateRoleOutputDTO, err error)
	DeleteRole(ctx context.Context, input request.RequestDeleteRole) (output usecase.DeleteRoleOutputDTO, err error)
}
