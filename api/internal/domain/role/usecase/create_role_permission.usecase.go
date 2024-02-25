package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionUseCase struct {
	repo repository.RolePermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase(
	repo repository.RolePermissionRepositoryInterface,
) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		repo: repo,
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (err error) {

	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.repo.CreateRolePermission(rolePermission)
	if err != nil {
		return
	}
	return
}
