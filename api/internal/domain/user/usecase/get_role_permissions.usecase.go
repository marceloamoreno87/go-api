package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRolePermissionsInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type GetRolePermissionsOutputDTO struct {
	ID            int32   `json:"id"`
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type GetRolePermissionsUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewGetRolePermissionsUseCase(DB config.SQLCInterface) *GetRolePermissionsUseCase {
	return &GetRolePermissionsUseCase{
		repo: repository.NewRolePermissionRepository(DB),
	}
}

func (uc *GetRolePermissionsUseCase) Execute(input GetRolePermissionsInputDTO) (output []GetRolePermissionsOutputDTO, err error) {
	rp, err := uc.repo.GetRolePermissionsByRole(input.RoleID)
	if err != nil {
		return
	}

	for _, v := range rp {
		output = append(output, GetRolePermissionsOutputDTO{
			ID:            v.GetID(),
			RoleID:        v.GetRoleID(),
			PermissionIDs: v.GetPermissionIDs(),
		})
	}
	return
}
