package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type UpdateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type UpdateRolePermissionUseCase struct {
	repo repository.RolePermissionRepositoryInterface
}

func NewUpdateRolePermissionUseCase(
	repo repository.RolePermissionRepositoryInterface,
) *UpdateRolePermissionUseCase {
	return &UpdateRolePermissionUseCase{
		repo: repo,
	}
}

func (uc *UpdateRolePermissionUseCase) Execute(input UpdateRolePermissionInputDTO) (err error) {
	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.repo.UpdateRolePermission(rolePermission, input.RoleID)
	if err != nil {
		return
	}
	return
}
