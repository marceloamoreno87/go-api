package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type Permissionrepository interface {
	CreatePermission(ctx context.Context, permission *entity.Permission) (err error)
	GetPermission(ctx context.Context, id int32) (output *entity.Permission, err error)
	GetPermissions(ctx context.Context, limit int32, offset int32) (output []*entity.Permission, err error)
	UpdatePermission(ctx context.Context, permission *entity.Permission, id int32) (err error)
	DeletePermission(ctx context.Context, id int32) (err error)
	GetPermissionByInternalName(ctx context.Context, internalName string) (output *entity.Permission, err error)
}

type PermissionRepository struct {
	db config.SQLCInterface
}

func NewPermissionRepository(db config.SQLCInterface) *PermissionRepository {
	return &PermissionRepository{
		db: db,
	}
}

func (repo *PermissionRepository) CreatePermission(ctx context.Context, permission *entity.Permission) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreatePermission(ctx, db.CreatePermissionParams{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
}

func (repo *PermissionRepository) GetPermission(ctx context.Context, id int32) (output *entity.Permission, err error) {
	p, err := repo.db.GetDbQueries().GetPermission(ctx, id)
	if err != nil {
		return
	}
	output = &entity.Permission{
		ID:           p.ID,
		Name:         p.Name,
		InternalName: p.InternalName,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
	return
}

func (repo *PermissionRepository) GetPermissions(ctx context.Context, limit int32, offset int32) (output []*entity.Permission, err error) {
	p, err := repo.db.GetDbQueries().GetPermissions(ctx, db.GetPermissionsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, permission := range p {
		output = append(output, &entity.Permission{
			ID:           permission.ID,
			Name:         permission.Name,
			InternalName: permission.InternalName,
			Description:  permission.Description,
			CreatedAt:    permission.CreatedAt,
			UpdatedAt:    permission.UpdatedAt,
		})
	}
	return
}

func (repo *PermissionRepository) UpdatePermission(ctx context.Context, permission *entity.Permission, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdatePermission(ctx, db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
}

func (repo *PermissionRepository) DeletePermission(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeletePermission(ctx, id)
}

func (repo *PermissionRepository) GetPermissionByInternalName(ctx context.Context, internalName string) (output *entity.Permission, err error) {
	p, err := repo.db.GetDbQueries().GetPermissionByInternalName(ctx, internalName)
	if err != nil {
		return
	}
	output = &entity.Permission{
		ID:           p.ID,
		Name:         p.Name,
		InternalName: p.InternalName,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
	return
}
