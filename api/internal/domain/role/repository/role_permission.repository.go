package repository

import (
	"context"
	"database/sql"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
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

func (r *RolePermissionRepository) GetRolePermissions(rolePermission *RoleEntity.RolePermission) (rolePermissions *RoleEntity.RolePermission, err error) {
	repo, err := r.DBQueries.GetRolePermissions(context.Background(), rolePermission.RoleId)
	if err != nil {
		return nil, err
	}

	rolePermissions = &RoleEntity.RolePermission{
		Role: &entity.Role{
			ID:           repo[0].RoleID,
			Name:         repo[0].Name,
			Description:  repo[0].Description,
			InternalName: repo[0].InternalName,
			CreatedAt:    repo[0].CreatedAt,
			UpdatedAt:    repo[0].UpdatedAt,
		},
	}

	rolePermissions.Permissions = make([]*PermissionEntity.Permission, 0)
	for _, p := range repo {
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

func (r *RolePermissionRepository) CreateRolePermission(rolePermission *entity.RolePermission) (rolePermissions *entity.RolePermission, err error) {
	return
}

func (r *RolePermissionRepository) UpdateRolePermission(rolePermission *entity.RolePermission) (rolePermissions *entity.RolePermission, err error) {
	return
}
