package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionServiceInterface interface {
	GetPermission(input request.RequestGetPermissionInputDTO) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(input request.RequestGetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error)
	CreatePermission(input request.RequestCreatePermissionInputDTO) (output usecase.CreatePermissionOutputDTO, err error)
	UpdatePermission(input request.RequestUpdatePermissionInputDTO) (output usecase.UpdatePermissionOutputDTO, err error)
	DeletePermission(input request.RequestDeletePermissionInputDTO) (output usecase.DeletePermissionOutputDTO, err error)
}
