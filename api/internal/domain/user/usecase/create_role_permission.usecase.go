package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionOutputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionUseCase struct {
	repo repository.RolePermissionrepository
}

func NewCreateRolePermissionUseCase(db config.SQLCInterface) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(db),
	}
}

func (uc *CreateRolePermissionUseCase) Execute(ctx context.Context, input CreateRolePermissionInputDTO) (output CreateRolePermissionOutputDTO, err error) {
	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.repo.CreateRolePermission(ctx, rolePermission)

	output = CreateRolePermissionOutputDTO{
		RoleID:        rolePermission.GetRoleID(),
		PermissionIDs: rolePermission.GetPermissionIDs(),
	}

	return
}
