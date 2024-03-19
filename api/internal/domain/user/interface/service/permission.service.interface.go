package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionServiceInterface interface {
	GetPermission(input service.RequestGetPermissionInputDTO) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(input service.RequestGetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error)
	CreatePermission(input service.RequestCreatePermissionInputDTO) (output usecase.CreatePermissionOutputDTO, err error)
	UpdatePermission(input service.RequestUpdatePermissionInputDTO) (output usecase.UpdatePermissionOutputDTO, err error)
	DeletePermission(input service.RequestDeletePermissionInputDTO) (output usecase.DeletePermissionOutputDTO, err error)
}
