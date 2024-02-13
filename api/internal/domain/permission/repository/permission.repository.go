package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type PermissionRepository struct {
	DBConn    *sql.DB
	DBQueries db.Querier
}

func NewPermissionRepository(DBConn *sql.DB) *PermissionRepository {
	return &PermissionRepository{
		DBConn:    DBConn,
		DBQueries: db.New(DBConn),
	}
}

func (repo *PermissionRepository) CreatePermission(permission *entity.Permission) (err error) {
	err = repo.DBQueries.CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) GetPermission(id int32) (permission *entity.Permission, err error) {
	p, err := repo.DBQueries.GetPermission(context.Background(), id)
	if err != nil {
		return &entity.Permission{}, err
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
	ps, err := repo.DBQueries.GetPermissions(context.Background(), db.GetPermissionsParams{
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
	err = repo.DBQueries.UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) DeletePermission(id int32) (err error) {
	err = repo.DBQueries.DeletePermission(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *PermissionRepository) GetPermissionByInternalName(internal_name string) (permission *entity.Permission, err error) {
	p, err := repo.DBQueries.GetPermissionByInternalName(context.Background(), internal_name)
	if err != nil {
		return &entity.Permission{}, err
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
