package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
	"github.com/marceloamoreno/goapi/internal/shared/repository"
)

type RoleRepositoryInterface interface {
	CreateRole(role *entity.Role) (err error)
	GetRole(id int32) (*entity.Role, error)
	GetRoleByInternalName(internal_name string) (*entity.Role, error)
	GetRoles(limit int32, offset int32) ([]*entity.Role, error)
	UpdateRole(role *entity.Role, id int32) (err error)
	DeleteRole(id int32) (err error)
	repository.RepositoryInterface
}

type RoleRepository struct {
	repository.Repository
}

func NewRoleRepository(DB config.DatabaseInterface) *RoleRepository {
	return &RoleRepository{
		Repository: *repository.NewRepository(DB),
	}
}

func (repo *RoleRepository) CreateRole(role *entity.Role) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) GetRole(id int32) (role *entity.Role, err error) {
	r, err := repo.GetDbQueries().GetRole(context.Background(), id)
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
	r, err := repo.GetDbQueries().GetRoleByInternalName(context.Background(), internal_name)
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
	rs, err := repo.GetDbQueries().GetRoles(context.Background(), db.GetRolesParams{
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
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           id,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	})
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRole(context.Background(), id)
	return
}
