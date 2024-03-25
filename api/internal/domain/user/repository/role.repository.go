package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type Rolerepository interface {
	CreateRole(ctx context.Context, role *entity.Role) (err error)
	GetRole(ctx context.Context, id int32) (output *entity.Role, err error)
	GetRoleByInternalName(ctx context.Context, internalName string) (output *entity.Role, err error)
	GetRoles(ctx context.Context, limit int32, offset int32) (output []*entity.Role, err error)
	UpdateRole(ctx context.Context, role *entity.Role, id int32) (err error)
	DeleteRole(ctx context.Context, id int32) (err error)
}

type RoleRepository struct {
	db config.SQLCInterface
}

func NewRoleRepository(db config.SQLCInterface) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (repo *RoleRepository) CreateRole(ctx context.Context, role *entity.Role) (err error) {

	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateRole(ctx, db.CreateRoleParams{
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
}

func (repo *RoleRepository) GetRole(ctx context.Context, id int32) (output *entity.Role, err error) {

	r, err := repo.db.GetDbQueries().GetRole(ctx, id)
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

func (repo *RoleRepository) GetRoleByInternalName(ctx context.Context, internalName string) (output *entity.Role, err error) {

	r, err := repo.db.GetDbQueries().GetRoleByInternalName(ctx, internalName)
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

func (repo *RoleRepository) GetRoles(ctx context.Context, limit int32, offset int32) (output []*entity.Role, err error) {

	r, err := repo.db.GetDbQueries().GetRoles(ctx, db.GetRolesParams{
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

func (repo *RoleRepository) UpdateRole(ctx context.Context, role *entity.Role, id int32) (err error) {

	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateRole(ctx, db.UpdateRoleParams{
		ID:           id,
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	})
}

func (repo *RoleRepository) DeleteRole(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeleteRole(ctx, id)
}
