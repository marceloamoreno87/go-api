package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRolePermissionByRoleIDInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type DeleteRolePermissionByRoleIDOutputDTO struct {
	RoleID int32 `json:"role_id"`
}

type DeleteRolePermissionByRoleIDUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewDeleteRolePermissionByRoleIDUseCase() *DeleteRolePermissionByRoleIDUseCase {
	return &DeleteRolePermissionByRoleIDUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *DeleteRolePermissionByRoleIDUseCase) Execute(input DeleteRolePermissionByRoleIDInputDTO) (output DeleteRolePermissionByRoleIDOutputDTO, err error) {
	err = uc.repo.DeleteRolePermissionByRoleID(input.RoleID)
	if err != nil {
		return
	}

	output = DeleteRolePermissionByRoleIDOutputDTO{
		RoleID: input.RoleID,
	}

	return
}
