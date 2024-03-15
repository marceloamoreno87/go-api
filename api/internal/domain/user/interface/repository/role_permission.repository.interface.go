package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type RolePermissionRepositoryInterface interface {
	GetRolePermissionsByRole(id int32) (output []entityInterface.RolePermissionInterface, err error)
	CreateRolePermission(rolePermission entityInterface.RolePermissionInterface) (err error)
	DeleteRolePermission(id int32) (err error)
}
