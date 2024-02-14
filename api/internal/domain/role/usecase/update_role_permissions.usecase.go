package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type UpdateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type UpdateRolePermissionUseCase struct {
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface
}

func NewUpdateRolePermissionUseCase(
	rolePermissionRepository RoleRepository.RolePermissionRepositoryInterface,
) *UpdateRolePermissionUseCase {
	return &UpdateRolePermissionUseCase{
		RolePermissionRepository: rolePermissionRepository,
	}
}

func (uc *UpdateRolePermissionUseCase) Execute(input UpdateRolePermissionInputDTO) (err error) {
	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.RolePermissionRepository.UpdateRolePermission(rolePermission, input.RoleID)
	if err != nil {
		return
	}
	return
}
