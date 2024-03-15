package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	GetRole(id int32) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(limit int32, offset int32) (output []usecase.GetRolesOutputDTO, err error)
	CreateRole(body io.ReadCloser) (output usecase.CreateRoleOutputDTO, err error)
	UpdateRole(id int32, body io.ReadCloser) (output usecase.UpdateRoleOutputDTO, err error)
	DeleteRole(id int32) (output usecase.DeleteRoleOutputDTO, err error)
}
