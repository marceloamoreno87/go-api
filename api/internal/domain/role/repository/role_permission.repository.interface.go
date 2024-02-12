package repository

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
)

type RolePermissionRepositoryInterface interface {
	CreateRolePermission(rolePermission *entity.RolePermission) (*entity.RolePermission, error)
}
