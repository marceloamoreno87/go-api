package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type PermissionRepository struct {
	config.SQLCInterface
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (repo *PermissionRepository) CreatePermission(permission entityInterface.PermissionInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
	return
}

func (repo *PermissionRepository) GetPermission(id int32) (permission entityInterface.PermissionInterface, err error) {
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

func (repo *PermissionRepository) GetPermissions(limit int32, offset int32) (permissions []entityInterface.PermissionInterface, err error) {
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

func (repo *PermissionRepository) UpdatePermission(permission entityInterface.PermissionInterface, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
	return
}

func (repo *PermissionRepository) DeletePermission(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).DeletePermission(context.Background(), id)
}

func (repo *PermissionRepository) GetPermissionByInternalName(internalName string) (permission entityInterface.PermissionInterface, err error) {
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
