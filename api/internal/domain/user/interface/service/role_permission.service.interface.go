package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(input request.RequestGetRolePermissionInputDTO) (output []usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(input request.RequestCreateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
	DeleteRolePermissionByRoleID(input request.RequestDeleteRolePermissionByRoleIDInputDTO) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
	UpdateRolePermission(input request.RequestUpdateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
}
