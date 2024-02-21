package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/pkg/api"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type RoleRepository struct {
	api.DefaultRepository
	DBConn    *sql.DB
	DBQueries *db.Queries
}

func NewRoleRepository(DBConn *sql.DB) *RoleRepository {
	return &RoleRepository{
		DBConn:    DBConn,
		DBQueries: db.New(DBConn),
	}
}

func (repo *RoleRepository) CreateRole(role *entity.Role) (err error) {
	err = repo.DBQueries.WithTx(repo.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) GetRole(id int32) (role *entity.Role, err error) {
	r, err := repo.DBQueries.GetRole(context.Background(), id)
	if err != nil {
		return
	}
	return &entity.Role{
		ID:           r.ID,
		Name:         r.Name,
		InternalName: r.InternalName,
		Description:  r.Description,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}, nil
}

func (repo *RoleRepository) GetRoleByInternalName(internal_name string) (role *entity.Role, err error) {
	r, err := repo.DBQueries.GetRoleByInternalName(context.Background(), internal_name)
	if err != nil {
		return
	}
	return &entity.Role{
		ID:           r.ID,
		Name:         r.Name,
		InternalName: r.InternalName,
		Description:  r.Description,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}, nil
}

func (repo *RoleRepository) GetRoles(limit int32, offset int32) (roles []*entity.Role, err error) {
	rs, err := repo.DBQueries.GetRoles(context.Background(), db.GetRolesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, r := range rs {
		roles = append(roles, &entity.Role{
			ID:           r.ID,
			Name:         r.Name,
			InternalName: r.InternalName,
			Description:  r.Description,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		})
	}
	return
}

func (repo *RoleRepository) UpdateRole(role *entity.Role, id int32) (err error) {
	err = repo.DBQueries.WithTx(repo.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	err = repo.DBQueries.WithTx(repo.GetTx()).DeleteRole(context.Background(), id)
	return
}
