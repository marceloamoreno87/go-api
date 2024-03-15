package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionServiceInterface interface {
	GetPermission(id int32) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(limit int32, offset int32) (output []usecase.GetPermissionsOutputDTO, err error)
	CreatePermission(body io.ReadCloser) (output usecase.CreatePermissionOutputDTO, err error)
	UpdatePermission(id int32, body io.ReadCloser) (output usecase.UpdatePermissionOutputDTO, err error)
	DeletePermission(id int32) (output usecase.DeletePermissionOutputDTO, err error)
}
