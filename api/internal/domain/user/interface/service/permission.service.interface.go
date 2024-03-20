package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionServiceInterface interface {
	GetPermission(input request.RequestGetPermission) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(input request.RequestGetPermissions) (output []usecase.GetPermissionsOutputDTO, err error)
	CreatePermission(input request.RequestCreatePermission) (output usecase.CreatePermissionOutputDTO, err error)
	UpdatePermission(input request.RequestUpdatePermission) (output usecase.UpdatePermissionOutputDTO, err error)
	DeletePermission(input request.RequestDeletePermission) (output usecase.DeletePermissionOutputDTO, err error)
}
