package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
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

type PermissionRepository struct {
	repository.Repository
}

func NewPermissionRepository(DB config.DatabaseInterface) *PermissionRepository {
	return &PermissionRepository{
		Repository: *repository.NewRepository(DB),
	}
}

func (repo *PermissionRepository) CreatePermission(permission *entity.Permission) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	return
}

func (repo *PermissionRepository) GetPermission(id int32) (permission *entity.Permission, err error) {
	p, err := repo.Repository.GetDbQueries().GetPermission(context.Background(), id)
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
	ps, err := repo.Repository.GetDbQueries().GetPermissions(context.Background(), db.GetPermissionsParams{
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
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	return
}

func (repo *PermissionRepository) DeletePermission(id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).DeletePermission(context.Background(), id)
	return
}

func (repo *PermissionRepository) GetPermissionByInternalName(internal_name string) (permission *entity.Permission, err error) {
	p, err := repo.Repository.GetDbQueries().GetPermissionByInternalName(context.Background(), internal_name)
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
