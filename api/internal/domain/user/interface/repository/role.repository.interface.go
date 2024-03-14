package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type RoleRepositoryInterface interface {
	CreateRole(role entityInterface.RoleInterface) (output entityInterface.RoleInterface, err error)
	GetRole(id int32) (output entityInterface.RoleInterface, err error)
	GetRoleByInternalName(internalName string) (output entityInterface.RoleInterface, err error)
	GetRoles(limit int32, offset int32) (output []entityInterface.RoleInterface, err error)
	UpdateRole(role entityInterface.RoleInterface, id int32) (output entityInterface.RoleInterface, err error)
	DeleteRole(id int32) (output entityInterface.RoleInterface, err error)
	config.SQLCInterface
}
