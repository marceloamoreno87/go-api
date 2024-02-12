package repository

import "github.com/marceloamoreno/goapi/internal/domain/role/entity"

type RolePermissionRepositoryInterface interface {
	GetRolePermissions(rolePermission *entity.RolePermission) (rolePermissions *entity.RolePermission, err error)
	CreateRolePermission(rolePermission *entity.RolePermission) (rolePermissions *entity.RolePermission, err error)
	UpdateRolePermission(rolePermission *entity.RolePermission) (rolePermissions *entity.RolePermission, err error)
}
