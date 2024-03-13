package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type PermissionRepositoryInterface interface {
	CreatePermission(permission *entity.Permission) (err error)
	GetPermission(id int32) (*entity.Permission, error)
	GetPermissions(limit int32, offset int32) (permissions []*entity.Permission, err error)
	UpdatePermission(permission *entity.Permission, id int32) (err error)
	DeletePermission(id int32) (err error)
	GetPermissionByInternalName(internalName string) (permission *entity.Permission, err error)
	config.SQLCInterface
}

type PermissionRepository struct {
	config.SQLCInterface
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (repo *PermissionRepository) CreatePermission(permission *entity.Permission) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	return
}

func (repo *PermissionRepository) GetPermission(id int32) (permission *entity.Permission, err error) {
	p, err := repo.GetDbQueries().GetPermission(context.Background(), id)
	if err != nil {
		return
	}
	return &entity.Permission{
		ID:           p.ID,
		Name:         p.Name,
		InternalName: p.InternalName,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}, nil
}

func (repo *PermissionRepository) GetPermissions(limit int32, offset int32) (permissions []*entity.Permission, err error) {
	ps, err := repo.GetDbQueries().GetPermissions(context.Background(), db.GetPermissionsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, p := range ps {
		permissions = append(permissions, &entity.Permission{
			ID:           p.ID,
			Name:         p.Name,
			InternalName: p.InternalName,
			Description:  p.Description,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		})
	}
	return
}

func (repo *PermissionRepository) UpdatePermission(permission *entity.Permission, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	return
}

func (repo *PermissionRepository) DeletePermission(id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).DeletePermission(context.Background(), id)
	return
}

func (repo *PermissionRepository) GetPermissionByInternalName(internalName string) (permission *entity.Permission, err error) {
	p, err := repo.GetDbQueries().GetPermissionByInternalName(context.Background(), internalName)
	if err != nil {
		return
	}
	return &entity.Permission{
		ID:           p.ID,
		Name:         p.Name,
		InternalName: p.InternalName,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}, nil
}
