package repository

import (
	"context"
	"database/sql"

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
func (rr *RolePermissionRepository) CreateRolePermission(role *entity.RolePermission) (*entity.RolePermission, error) {
	repo, err := rr.DBQueries.CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
		RoleID:       role.RoleId,
		PermissionID: role.PermissionId,
	})
	if err != nil {
		return &entity.RolePermission{}, err
	}

	return &entity.RolePermission{
		RoleId:       repo.RoleID,
		PermissionId: repo.PermissionID,
		Role:         role.GetRole(),
		Permission:   role.GetPermission(),
	}, nil
}
