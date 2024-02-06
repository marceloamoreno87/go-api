package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
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

func (rr *RoleRepository) CreateRole(role *entity.Role) (*entity.Role, error) {
	repo, err := rr.DBQueries.CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	if err != nil {
		return &entity.Role{}, err
	}

	return &entity.Role{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (rr *RoleRepository) GetRole(id int32) (*entity.Role, error) {

	repo, err := rr.DBQueries.GetRole(context.Background(), id)
	if err != nil {
		return &entity.Role{}, err
	}

	return &entity.Role{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (rr *RoleRepository) GetRoleByInternalName(internal_name string) (*entity.Role, error) {
	repo, err := rr.DBQueries.GetRoleByInternalName(context.Background(), internal_name)
	if err != nil {
		return &entity.Role{}, err
	}

	return &entity.Role{
		ID:           repo.ID,
		Name:         repo.Name,
		InternalName: repo.InternalName,
		Description:  repo.Description,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
	}, nil
}

func (rr *RoleRepository) GetRoles(limit int32, offset int32) (roles []*entity.Role, err error) {
	repo, err := rr.DBQueries.GetRoles(context.Background(), db.GetRolesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []*entity.Role{}, err
	}

	for _, r := range repo {
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

func (rr *RoleRepository) UpdateRole(role *entity.Role, id int32) (*entity.Role, error) {

	r, err := rr.DBQueries.UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           id,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
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

func (ur *RoleRepository) DeleteRole(id int32) (*entity.Role, error) {
	r, err := ur.DBQueries.DeleteRole(context.Background(), id)

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
