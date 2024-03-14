package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type UpdateRolePermissionUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewUpdateRolePermissionUseCase() *UpdateRolePermissionUseCase {
	return &UpdateRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *UpdateRolePermissionUseCase) Execute(input UpdateRolePermissionInputDTO) (err error) {
	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	if err = uc.repo.UpdateRolePermission(rolePermission, input.RoleID); err != nil {
		return
	}
	return
}
