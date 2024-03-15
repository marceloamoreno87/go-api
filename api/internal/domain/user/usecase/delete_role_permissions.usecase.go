package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type DeleteRolePermissionOutputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type DeleteRolePermissionUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewDeleteRolePermissionUseCase() *DeleteRolePermissionUseCase {
	return &DeleteRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *DeleteRolePermissionUseCase) Execute(input DeleteRolePermissionInputDTO) (output DeleteRolePermissionOutputDTO, err error) {
	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	rp, err := uc.repo.DeleteRolePermission(rolePermission, input.RoleID)
	if err != nil {
		return
	}

	output = DeleteRolePermissionOutputDTO{
		RoleID:        rp.GetRoleID(),
		PermissionIDs: rp.GetPermissionIDs(),
	}

	return
}
