package entity

import (
	Permission "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
)

type RolePermissionInterface interface {
	Validate() (err error)
	GetRole() *Role
	GetPermission() *Permission.Permission
	SetRole(role *Role)
	SetPermission(permission *Permission.Permission)
	GetRoleID()
	GetPermissionID()
	SetRoleID(roleId int32)
	SetPermissionID(permissionId int32)
}
