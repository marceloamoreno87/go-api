package repository

import "github.com/marceloamoreno/goapi/internal/domain/permission/entity"

type PermissionRepositoryInterface interface {
	CreatePermission(permission *entity.Permission) (*entity.Permission, error)
	GetPermission(id int32) (*entity.Permission, error)
	GetPermissions(limit int32, offset int32) ([]*entity.Permission, error)
	UpdatePermission(permission *entity.Permission) (*entity.Permission, error)
	DeletePermission(id int32) (*entity.Permission, error)
	GetPermissionByInternalName(internalName string) (*entity.Permission, error)
}
