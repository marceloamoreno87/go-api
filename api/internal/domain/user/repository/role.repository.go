package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RoleRepository struct {
	config.SQLCInterface
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (repo *RoleRepository) CreateRole(role entityInterface.RoleInterface) (output db.Role, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateRole(context.Background(), db.CreateRoleParams{
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) GetRole(id int32) (output db.Role, err error) {
	output, err = repo.GetDbQueries().GetRole(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) GetRoleByInternalName(internalName string) (output db.Role, err error) {
	output, err = repo.GetDbQueries().GetRoleByInternalName(context.Background(), internalName)
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) GetRoles(limit int32, offset int32) (output []db.Role, err error) {
	output, err = repo.GetDbQueries().GetRoles(context.Background(), db.GetRolesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) UpdateRole(role entityInterface.RoleInterface, id int32) (output db.Role, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateRole(context.Background(), db.UpdateRoleParams{
		ID:           id,
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *RoleRepository) DeleteRole(id int32) (output db.Role, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRole(context.Background(), id)
	if err != nil {
		return
	}
	return
}
