package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type CreateRolePermissionUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase() *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (err error) {

	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	if err = uc.repo.CreateRolePermission(rolePermission); err != nil {
		return
	}
	return
}
