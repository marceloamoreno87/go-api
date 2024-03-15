package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type PermissionRepository struct {
	DB config.SQLCInterface
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{
		DB: config.Sqcl,
	}
}

func (repo *PermissionRepository) CreatePermission(permission entityInterface.PermissionInterface) (err error) {
	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
}

func (repo *PermissionRepository) GetPermission(id int32) (output entityInterface.PermissionInterface, err error) {
	p, err := repo.DB.GetDbQueries().GetPermission(context.Background(), id)
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

func (repo *PermissionRepository) GetPermissions(limit int32, offset int32) (output []entityInterface.PermissionInterface, err error) {
	p, err := repo.DB.GetDbQueries().GetPermissions(context.Background(), db.GetPermissionsParams{
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

func (repo *PermissionRepository) UpdatePermission(permission entityInterface.PermissionInterface, id int32) (err error) {
	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
}

func (repo *PermissionRepository) DeletePermission(id int32) (err error) {

	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).DeletePermission(context.Background(), id)
}

func (repo *PermissionRepository) GetPermissionByInternalName(internalName string) (output entityInterface.PermissionInterface, err error) {
	p, err := repo.DB.GetDbQueries().GetPermissionByInternalName(context.Background(), internalName)
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
