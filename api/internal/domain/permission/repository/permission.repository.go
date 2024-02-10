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

func (pr *PermissionRepository) CreatePermission(permission *entity.Permission) (*entity.Permission, error) {
	repo, err := pr.DBQueries.CreatePermission(context.Background(), db.CreatePermissionParams{
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	if err != nil {
		return &entity.Permission{}, err
	}

	return &entity.Permission{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (pr *PermissionRepository) GetPermission(id int32) (*entity.Permission, error) {

	repo, err := pr.DBQueries.GetPermission(context.Background(), id)
	if err != nil {
		return &entity.Permission{}, err
	}

	return &entity.Permission{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (pr *PermissionRepository) GetPermissions(limit int32, offset int32) ([]*entity.Permission, error) {

	repo, err := pr.DBQueries.GetPermissions(context.Background(), db.GetPermissionsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []*entity.Permission{}, err
	}

	var permissions []*entity.Permission
	for _, p := range repo {
		permissions = append(permissions, &entity.Permission{
			ID:           p.ID,
			Name:         p.Name,
			InternalName: p.InternalName,
			Description:  p.Description,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		})
	}

	return permissions, nil
}

func (pr *PermissionRepository) UpdatePermission(permission *entity.Permission) (*entity.Permission, error) {
	repo, err := pr.DBQueries.UpdatePermission(context.Background(), db.UpdatePermissionParams{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	})
	if err != nil {
		return &entity.Permission{}, err
	}

	return &entity.Permission{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (pr *PermissionRepository) DeletePermission(id int32) (*entity.Permission, error) {
	repo, err := pr.DBQueries.DeletePermission(context.Background(), id)

	if err != nil {
		return &entity.Permission{}, err
	}
	return &entity.Permission{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (pr *PermissionRepository) GetPermissionByInternalName(name string) (*entity.Permission, error) {
	repo, err := pr.DBQueries.GetPermissionByInternalName(context.Background(), name)
	if err != nil {
		return &entity.Permission{}, err
	}

	return &entity.Permission{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}
