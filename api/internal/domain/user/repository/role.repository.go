package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RoleRepository struct {
	DB config.SQLCInterface
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		DB: config.Sqcl,
	}
}

func (repo *RoleRepository) CreateRole(role entityInterface.RoleInterface) (err error) {
	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
}

func (repo *RoleRepository) GetRole(id int32) (output entityInterface.RoleInterface, err error) {
	r, err := repo.DB.GetDbQueries().GetRole(context.Background(), id)
	if err != nil {
		return
	}
	output = &entity.Role{
		ID:           r.ID,
		Name:         r.Name,
		InternalName: r.InternalName,
		Description:  r.Description,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
	return
}

func (repo *RoleRepository) GetRoleByInternalName(internalName string) (output entityInterface.RoleInterface, err error) {
	r, err := repo.DB.GetDbQueries().GetRoleByInternalName(context.Background(), internalName)
	if err != nil {
		return
	}
	output = &entity.Role{
		ID:           r.ID,
		Name:         r.Name,
		InternalName: r.InternalName,
		Description:  r.Description,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
	return
}

func (repo *RoleRepository) GetRoles(limit int32, offset int32) (output []entityInterface.RoleInterface, err error) {
	r, err := repo.DB.GetDbQueries().GetRoles(context.Background(), db.GetRolesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, role := range r {
		output = append(output, &entity.Role{
			ID:           role.ID,
			Name:         role.Name,
			InternalName: role.InternalName,
			Description:  role.Description,
			CreatedAt:    role.CreatedAt,
			UpdatedAt:    role.UpdatedAt,
		})
	}
	return
}

func (repo *RoleRepository) UpdateRole(role entityInterface.RoleInterface, id int32) (err error) {
	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           id,
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
}

func (repo *RoleRepository) DeleteRole(id int32) (err error) {
	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).DeleteRole(context.Background(), id)
}
