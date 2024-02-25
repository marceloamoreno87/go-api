package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/infra/database"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type RoleRepository struct {
	database.Repository
}

func NewRoleRepository(dbConn *sql.DB) *RoleRepository {
	return &RoleRepository{
		Repository: *database.NewRepository(dbConn),
	}
}

func (repo *RoleRepository) CreateRole(role *entity.Role) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) GetRole(id int32) (role *entity.Role, err error) {
	r, err := repo.Repository.GetDbQueries().GetRole(context.Background(), id)
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
	r, err := repo.Repository.GetDbQueries().GetRoleByInternalName(context.Background(), internal_name)
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
	rs, err := repo.Repository.GetDbQueries().GetRoles(context.Background(), db.GetRolesParams{
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
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).DeleteRole(context.Background(), id)
	return
}
