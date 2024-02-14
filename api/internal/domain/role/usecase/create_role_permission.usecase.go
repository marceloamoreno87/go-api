package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionUseCase struct {
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase(
	rolePermissionRepository RoleRepository.RolePermissionRepositoryInterface,
) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		RolePermissionRepository: rolePermissionRepository,
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (err error) {

	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.RolePermissionRepository.CreateRolePermission(rolePermission)
	if err != nil {
		return
	}
	return
}
