package repository

import (
	"context"
	"database/sql"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RolePermissionEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
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

func (repo *RolePermissionRepository) GetRolePermissions(id int32) (rolePermissions *RolePermissionEntity.RolePermission, err error) {
	rps, err := repo.DBQueries.GetRolePermissions(context.Background(), id)
	if err != nil {
		return
	}
	rolePermissions = &RolePermissionEntity.RolePermission{
		Role: &RoleEntity.Role{
			ID:           rps[0].RoleID,
			Name:         rps[0].Name,
			InternalName: rps[0].InternalName,
			Description:  rps[0].Description,
			CreatedAt:    rps[0].CreatedAt,
			UpdatedAt:    rps[0].UpdatedAt,
		},
	}

	for _, rp := range rps {
		rolePermissions.Permissions = append(rolePermissions.Permissions, &PermissionEntity.Permission{
			ID:           rp.PermissionID,
			Name:         rp.Name_2,
			InternalName: rp.InternalName_2,
			Description:  rp.Description_2,
			CreatedAt:    rp.CreatedAt_2,
			UpdatedAt:    rp.UpdatedAt_2,
		})
	}
	return
}

func (repo *RolePermissionRepository) CreateRolePermission(rolePermission *RolePermissionEntity.RolePermission) (err error) {
	for _, id := range rolePermission.PermissionIDs {
		err = repo.DBQueries.CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
			RoleID:       rolePermission.RoleID,
			PermissionID: id,
		})
		if err != nil {
			return
		}
	}
	return
}

func (repo *RolePermissionRepository) UpdateRolePermission(rolePermission *entity.RolePermission, id int32) (err error) {
	err = repo.DBQueries.DeleteRolePermission(context.Background(), id)
	if err != nil {
		return
	}
	for _, permissionId := range rolePermission.PermissionIDs {
		err = repo.DBQueries.CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
			RoleID:       rolePermission.RoleID,
			PermissionID: permissionId,
		})
		if err != nil {
			return
		}
	}
	return

}
