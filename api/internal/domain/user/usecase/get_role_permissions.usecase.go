package usecase

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRolePermissionsInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type GetRolePermissionsOutputDTO struct {
	Role        entityInterface.RoleInterface         `json:"role"`
	Permissions []entityInterface.PermissionInterface `json:"permissions"`
}

type GetRolePermissionsUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewGetRolePermissionsUseCase() *GetRolePermissionsUseCase {
	return &GetRolePermissionsUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *GetRolePermissionsUseCase) Execute(input GetRolePermissionsInputDTO) (output GetRolePermissionsOutputDTO, err error) {
	rolePermission, err := uc.repo.GetRolePermissionsByRole(input.RoleID)
	if err != nil {
		return
	}

	rolePermission.SetRoleID(input.RoleID)

	output.Role = rolePermission.GetRole()
	output.Permissions = rolePermission.GetPermissions()

	return
}
