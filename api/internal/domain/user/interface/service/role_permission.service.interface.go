package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(input service.RequestGetRolePermissionInputDTO) (output []usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(input service.RequestCreateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
	DeleteRolePermissionByRoleID(input service.RequestDeleteRolePermissionByRoleIDInputDTO) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
	UpdateRolePermission(input service.RequestUpdateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
}
