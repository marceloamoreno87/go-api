package usecase

import (
	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type GetRolePermissionsInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type GetRolePermissionsOutputDTO struct {
	Role        *RoleEntity.Role               `json:"role"`
	Permissions []*PermissionEntity.Permission `json:"permissions"`
}

type GetRolePermissionsUseCase struct {
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface
}

func NewGetRolePermissionsUseCase(
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface,
) *GetRolePermissionsUseCase {
	return &GetRolePermissionsUseCase{
		RolePermissionRepository: RolePermissionRepository,
	}
}

func (uc *GetRolePermissionsUseCase) Execute(input GetRolePermissionsInputDTO) (output GetRolePermissionsOutputDTO, err error) {
	rolePermission := &RoleEntity.RolePermission{
		RoleID: input.RoleID,
	}
	rolePermission, err = uc.RolePermissionRepository.GetRolePermissions(input.RoleID)
	if err != nil {
		return
	}

	output.Role = rolePermission.Role
	output.Permissions = rolePermission.Permissions

	return
}
