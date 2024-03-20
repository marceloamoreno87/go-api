package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(input request.RequestGetRolePermission) (output []usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(input request.RequestCreateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error)
	DeleteRolePermissionByRoleID(input request.RequestDeleteRolePermissionByRoleID) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
	UpdateRolePermission(input request.RequestUpdateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error)
}
