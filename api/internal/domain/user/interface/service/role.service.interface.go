package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	GetRole(input service.RequestGetRoleInputDTO) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(input service.RequestGetRolesInputDTO) (output []usecase.GetRolesOutputDTO, err error)
	CreateRole(input service.RequestCreateRoleInputDTO) (output usecase.CreateRoleOutputDTO, err error)
	UpdateRole(input service.RequestUpdateRoleInputDTO) (output usecase.UpdateRoleOutputDTO, err error)
	DeleteRole(input service.RequestDeleteRoleInputDTO) (output usecase.DeleteRoleOutputDTO, err error)
}
