package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionOutputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase(DB config.SQLCInterface) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(DB),
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (output CreateRolePermissionOutputDTO, err error) {

	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	err = uc.repo.CreateRolePermission(rolePermission)

	output = CreateRolePermissionOutputDTO{
		RoleID:        rolePermission.GetRoleID(),
		PermissionIDs: rolePermission.GetPermissionIDs(),
	}

	return
}
