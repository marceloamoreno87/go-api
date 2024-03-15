package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(id int32) (output []usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(body io.ReadCloser) (output usecase.CreateRolePermissionOutputDTO, err error)
	DeleteRolePermission(id int32, body io.ReadCloser) (output usecase.DeleteRolePermissionOutputDTO, err error)
}
