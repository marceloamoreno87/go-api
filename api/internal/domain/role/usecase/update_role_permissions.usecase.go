package usecase

import (
	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	PermissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RolePermissionEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type UpdateRolePermissionInputDTO struct {
	RoleID       int32   `json:"role_id"`
	PermissionID []int32 `json:"permission_id"`
}

type UpdateRolePermissionOutputDTO struct {
	RoleID       int32                          `json:"role_id"`
	PermissionID []int32                        `json:"permission_id"`
	Role         RoleEntity.Role                `json:"role"`
	Permission   []*PermissionEntity.Permission `json:"permission"`
}

type UpdateRolePermissionUseCase struct {
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface
	PermissionReposity       PermissionRepository.PermissionRepositoryInterface
}

func NewUpdateRolePermissionUseCase(
	permissionRepository PermissionRepository.PermissionRepositoryInterface,
) *UpdateRolePermissionUseCase {
	return &UpdateRolePermissionUseCase{
		PermissionReposity: permissionRepository,
	}
}

func (uc *UpdateRolePermissionUseCase) Execute(input UpdateRolePermissionInputDTO) (output UpdateRolePermissionOutputDTO, err error) {
	rolePermission, err := RolePermissionEntity.NewRolePermission(&RoleEntity.Role{ID: input.RoleID}, []*PermissionEntity.Permission{})
	if err != nil {
		return UpdateRolePermissionOutputDTO{}, err
	}
	rolePermission, err = uc.RolePermissionRepository.UpdateRolePermission(rolePermission)
	if err != nil {
		return UpdateRolePermissionOutputDTO{}, err
	}

	output = UpdateRolePermissionOutputDTO{
		RoleID:       rolePermission.Role.ID,
		PermissionID: input.PermissionID,
		Role:         *rolePermission.Role,
		Permission:   rolePermission.Permissions,
	}

	return
}
