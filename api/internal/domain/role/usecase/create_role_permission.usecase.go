package usecase

import (
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID       int32   `json:"role_id"`
	PermissionID []int32 `json:"permission_id"`
}

type CreateRolePermissionUseCase struct {
	RoleRepository RoleRepository.RoleRepositoryInterface
}

func NewCreateRolePermissionUseCase(
	roleRepository RoleRepository.RoleRepositoryInterface,
) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		RoleRepository: roleRepository,
	}
}
