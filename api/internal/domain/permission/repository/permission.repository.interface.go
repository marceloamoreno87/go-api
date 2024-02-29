package repository

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/shared/repository"
)

type PermissionRepositoryInterface interface {
	CreatePermission(permission *entity.Permission) (err error)
	GetPermission(id int32) (*entity.Permission, error)
	GetPermissions(limit int32, offset int32) (permissions []*entity.Permission, err error)
	UpdatePermission(permission *entity.Permission, id int32) (err error)
	DeletePermission(id int32) (err error)
	GetPermissionByInternalName(internal_name string) (permission *entity.Permission, err error)
	repository.RepositoryInterface
}
