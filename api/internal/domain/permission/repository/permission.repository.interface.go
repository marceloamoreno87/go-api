package repository

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type PermissionRepositoryInterface interface {
	CreatePermission(permission *entity.Permission) (err error)
	GetPermission(id int32) (*entity.Permission, error)
	GetPermissions(page *api.Paginate) ([]*entity.Permission, error)
	UpdatePermission(permission *entity.Permission) (err error)
	DeletePermission(id int32) (err error)
	GetPermissionByInternalName(internal_name string) (*entity.Permission, error)
}
