package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RoleRepository struct {
	config.SQLCInterface
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (repo *RoleRepository) CreateRole(role entityInterface.RoleInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
	return
}

func (repo *RoleRepository) GetRole(id int32) (role entityInterface.RoleInterface, err error) {
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

func (repo *RoleRepository) GetRoleByInternalName(internalName string) (role entityInterface.RoleInterface, err error) {
	r, err := repo.GetDbQueries().GetRoleByInternalName(context.Background(), internalName)
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

func (repo *RoleRepository) GetRoles(limit int32, offset int32) (roles []entityInterface.RoleInterface, err error) {
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

func (repo *RoleRepository) UpdateRole(role entityInterface.RoleInterface, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           id,
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRole(context.Background(), id)
}
