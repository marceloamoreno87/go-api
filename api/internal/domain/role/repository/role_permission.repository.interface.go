package repository

import "github.com/marceloamoreno/goapi/internal/domain/role/entity"

type RolePermissionRepositoryInterface interface {
	GetRolePermissions(id int32) (rolePermissions *entity.RolePermission, err error)
	CreateRolePermission(rolePermission *entity.RolePermission) (err error)
	UpdateRolePermission(rolePermission *entity.RolePermission, id int32) (err error)
}
