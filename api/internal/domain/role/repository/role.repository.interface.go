package repository

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type RoleRepositoryInterface interface {
	CreateRole(role *entity.Role) (err error)
	GetRole(id int32) (*entity.Role, error)
	GetRoleByInternalName(internal_name string) (*entity.Role, error)
	GetRoles(page *api.Paginate) ([]*entity.Role, error)
	UpdateRole(role *entity.Role, id int32) (err error)
	DeleteRole(id int32) (err error)
}
