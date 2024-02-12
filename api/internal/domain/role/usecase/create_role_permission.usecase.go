package usecase

import (
	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	PermissionRepository "github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	PermissionUsecase "github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleRepository "github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

type CreateRolePermissionOutputDTO struct {
	RoleID       int32                       `json:"role_id"`
	PermissionID int32                       `json:"permission_id"`
	Role         RoleEntity.Role             `json:"role"`
	Permission   PermissionEntity.Permission `json:"permission"`
}

type CreateRolePermissionUseCase struct {
	RolePermissionRepository RoleRepository.RolePermissionRepositoryInterface
	RoleRepository           RoleRepository.RoleRepositoryInterface
	PermissionReposity       PermissionRepository.PermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase(
	roleRepository RoleRepository.RoleRepositoryInterface,
	permissionRepository PermissionRepository.PermissionRepositoryInterface,
) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		RoleRepository:     roleRepository,
		PermissionReposity: permissionRepository,
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (output CreateRolePermissionOutputDTO, err error) {
	ru := NewGetRoleUseCase(uc.RoleRepository)
	inputRoleDTO := GetRoleInputDTO{
		ID: input.RoleID,
	}
	role, err := ru.Execute(inputRoleDTO)

	if err != nil {
		return CreateRolePermissionOutputDTO{}, err
	}

	roleEntity := RoleEntity.Role{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
	}

	pu := PermissionUsecase.NewGetPermissionUseCase(uc.PermissionReposity)
	inputPermissionDTO := PermissionUsecase.GetPermissionInputDTO{
		ID: input.PermissionID,
	}

	permission, err := pu.Execute(inputPermissionDTO)
	if err != nil {
		return CreateRolePermissionOutputDTO{}, err
	}

	permissionEntity := PermissionEntity.Permission{
		ID:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}

	rolePermission, err := entity.NewRolePermission(&roleEntity, &permissionEntity)
	if err != nil {
		return CreateRolePermissionOutputDTO{}, err
	}

	u, err := uc.RolePermissionRepository.CreateRolePermission(rolePermission)
	if err != nil {
		return CreateRolePermissionOutputDTO{}, err
	}

	output = CreateRolePermissionOutputDTO{
		RoleID:       u.GetRoleID(),
		PermissionID: u.GetPermissionID(),
		Role:         *u.GetRole(),
		Permission:   *u.GetPermission(),
	}

	return
}
