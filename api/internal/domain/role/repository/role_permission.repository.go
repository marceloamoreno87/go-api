package repository

import (
	"context"
	"database/sql"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type RolePermissionRepository struct {
	DBConn    *sql.DB
	DBQueries db.Querier
}

func NewRolePermissionRepository(DBConn *sql.DB) *RolePermissionRepository {
	return &RolePermissionRepository{
		DBConn:    DBConn,
		DBQueries: db.New(DBConn),
	}
}

func (repo *RolePermissionRepository) GetRolePermissions(id int32) (rolePermissions *entity.RolePermission, err error) {
	rps, err := repo.DBQueries.GetRolePermissions(context.Background(), id)
	if err != nil {
		return nil, err
	}
	rolePermissions.Role = &entity.Role{
		ID:           rps[0].RoleID,
		Name:         rps[0].Name,
		Description:  rps[0].Description,
		InternalName: rps[0].InternalName,
		CreatedAt:    rps[0].CreatedAt,
		UpdatedAt:    rps[0].UpdatedAt,
	}

	// TODO: Refactor this
	for _, p := range rps {
		rolePermissions.Permissions = append(rolePermissions.Permissions, &PermissionEntity.Permission{
			ID:           p.PermissionID,
			Name:         p.Name_2,
			Description:  p.Description_2,
			InternalName: p.InternalName_2,
			CreatedAt:    p.CreatedAt_2,
			UpdatedAt:    p.UpdatedAt_2,
		})
	}
	return
}

func (repo *RolePermissionRepository) CreateRolePermission(rolePermission *entity.RolePermission) (err error) {
	for _, p := range rolePermission.Permissions {
		err = repo.DBQueries.CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
			RoleID:       rolePermission.RoleID,
			PermissionID: p.ID,
		})
		if err != nil {
			return
		}
	}
	return
}

func (repo *RolePermissionRepository) UpdateRolePermission(rolePermission *entity.RolePermission) (err error) {
	err = repo.DBQueries.DeleteRolePermission(context.Background(), rolePermission.RoleID)
	if err != nil {
		return
	}
	for _, p := range rolePermission.Permissions {
		err = repo.DBQueries.CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
			RoleID:       rolePermission.RoleID,
			PermissionID: p.ID,
		})
		if err != nil {
			return
		}
	}
	return
}
