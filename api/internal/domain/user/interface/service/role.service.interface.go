package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	GetRole(input request.RequestGetRoleInputDTO) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(input request.RequestGetRolesInputDTO) (output []usecase.GetRolesOutputDTO, err error)
	CreateRole(input request.RequestCreateRoleInputDTO) (output usecase.CreateRoleOutputDTO, err error)
	UpdateRole(input request.RequestUpdateRoleInputDTO) (output usecase.UpdateRoleOutputDTO, err error)
	DeleteRole(input request.RequestDeleteRoleInputDTO) (output usecase.DeleteRoleOutputDTO, err error)
}
