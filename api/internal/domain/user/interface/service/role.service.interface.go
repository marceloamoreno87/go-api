package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	GetRole(input request.RequestGetRole) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(input request.RequestGetRoles) (output []usecase.GetRolesOutputDTO, err error)
	CreateRole(input request.RequestCreateRole) (output usecase.CreateRoleOutputDTO, err error)
	UpdateRole(input request.RequestUpdateRole) (output usecase.UpdateRoleOutputDTO, err error)
	DeleteRole(input request.RequestDeleteRole) (output usecase.DeleteRoleOutputDTO, err error)
}
