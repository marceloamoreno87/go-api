package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRolePermissionInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type DeleteRolePermissionOutputDTO struct {
	RoleID int32 `json:"role_id"`
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
	err = uc.repo.DeleteRolePermission(input.RoleID)
	if err != nil {
		return
	}

	output = DeleteRolePermissionOutputDTO{
		RoleID: input.RoleID,
	}

	return
}
