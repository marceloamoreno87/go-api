package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type PermissionRepositoryInterface interface {
	CreatePermission(permission entityInterface.PermissionInterface) (output entityInterface.PermissionInterface, err error)
	GetPermission(id int32) (output entityInterface.PermissionInterface, err error)
	GetPermissions(limit int32, offset int32) (output []entityInterface.PermissionInterface, err error)
	UpdatePermission(permission entityInterface.PermissionInterface, id int32) (output entityInterface.PermissionInterface, err error)
	DeletePermission(id int32) (output entityInterface.PermissionInterface, err error)
	GetPermissionByInternalName(internalName string) (output entityInterface.PermissionInterface, err error)
}
