package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type PermissionRepository struct {
	config.SQLCInterface
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (repo *PermissionRepository) CreatePermission(permission entityInterface.PermissionInterface) (output db.Permission, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) GetPermission(id int32) (output db.Permission, err error) {
	output, err = repo.GetDbQueries().GetPermission(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) GetPermissions(limit int32, offset int32) (output []db.Permission, err error) {
	output, err = repo.GetDbQueries().GetPermissions(context.Background(), db.GetPermissionsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) UpdatePermission(permission entityInterface.PermissionInterface, id int32) (output db.Permission, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           id,
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	})
	return
}

func (repo *PermissionRepository) DeletePermission(id int32) (output db.Permission, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).DeletePermission(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) GetPermissionByInternalName(internalName string) (output db.Permission, err error) {
	output, err = repo.GetDbQueries().GetPermissionByInternalName(context.Background(), internalName)
	if err != nil {
		return
	}
	return
}
