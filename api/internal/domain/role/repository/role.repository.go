package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/pkg/api"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type RoleRepository struct {
	DBConn    *sql.DB
	DBQueries db.Querier
}

func NewRoleRepository(DBConn *sql.DB) *RoleRepository {
	return &RoleRepository{
		DBConn:    DBConn,
		DBQueries: db.New(DBConn),
	}
}

func (repo *RoleRepository) CreateRole(role *entity.Role) (err error) {
	err = repo.DBQueries.CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) GetRole(id int32) (*entity.Role, error) {
	r, err := repo.DBQueries.GetRole(context.Background(), id)
	if err != nil {
		return &entity.Role{}, err
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

func (repo *RoleRepository) GetRoleByInternalName(internal_name string) (*entity.Role, error) {
	r, err := repo.DBQueries.GetRoleByInternalName(context.Background(), internal_name)
	if err != nil {
		return &entity.Role{}, err
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

func (repo *RoleRepository) GetRoles(page *api.Paginate) (roles []*entity.Role, err error) {
	rs, err := repo.DBQueries.GetRoles(context.Background(), db.GetRolesParams{
		Limit:  page.Limit,
		Offset: page.Offset,
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
	err = repo.DBQueries.UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	err = repo.DBQueries.DeleteRole(context.Background(), id)
	if err != nil {
		return err
	}
	return
}
