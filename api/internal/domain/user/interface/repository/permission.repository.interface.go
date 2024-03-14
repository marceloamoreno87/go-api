package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type PermissionRepositoryInterface interface {
	CreatePermission(permission entityInterface.PermissionInterface) (err error)
	GetPermission(id int32) (entityInterface.PermissionInterface, error)
	GetPermissions(limit int32, offset int32) (permissions []entityInterface.PermissionInterface, err error)
	UpdatePermission(permission entityInterface.PermissionInterface, id int32) (err error)
	DeletePermission(id int32) (err error)
	GetPermissionByInternalName(internalName string) (permission entityInterface.PermissionInterface, err error)
	config.SQLCInterface
}
