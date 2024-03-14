package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type RoleRepositoryInterface interface {
	CreateRole(role entityInterface.RoleInterface) (err error)
	GetRole(id int32) (entityInterface.RoleInterface, error)
	GetRoleByInternalName(internalName string) (entityInterface.RoleInterface, error)
	GetRoles(limit int32, offset int32) ([]entityInterface.RoleInterface, error)
	UpdateRole(role entityInterface.RoleInterface, id int32) (err error)
	DeleteRole(id int32) (err error)
	config.SQLCInterface
}
