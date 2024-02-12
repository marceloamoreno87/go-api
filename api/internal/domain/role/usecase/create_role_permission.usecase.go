package usecase

import (
	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID       int32   `json:"role_id"`
	PermissionID []int32 `json:"permission_id"`
}

type CreateRolePermissionOutputDTO struct {
	RoleID       int32                          `json:"role_id"`
	PermissionID []int32                        `json:"permission_id"`
	Role         RoleEntity.Role                `json:"role"`
	Permission   []*PermissionEntity.Permission `json:"permission"`
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
