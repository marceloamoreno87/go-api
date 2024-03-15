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

type CreateRolePermissionOutputDTO struct {
	ID     int32 `json:"id"`
	RoleID int32 `json:"role_id"`
}

type CreateRolePermissionUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewCreateRolePermissionUseCase() *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *CreateRolePermissionUseCase) Execute(input CreateRolePermissionInputDTO) (output CreateRolePermissionOutputDTO, err error) {

	rolePermission, err := entity.NewRolePermission(input.RoleID, input.PermissionIDs)
	if err != nil {
		return
	}

	rp, err := uc.repo.CreateRolePermission(rolePermission)
	if err != nil {
		return
	}

	for _, p := range rp {
		output = CreateRolePermissionOutputDTO{
			ID:     p.GetID(),
			RoleID: p.GetRoleID(),
		}
	}

	return
}
