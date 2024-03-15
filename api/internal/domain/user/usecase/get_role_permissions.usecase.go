package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRolePermissionsInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type GetRolePermissionsOutputDTO struct {
	ID     int32 `json:"id"`
	RoleID int32 `json:"role_id"`
}

type GetRolePermissionsUseCase struct {
	repo repositoryInterface.RolePermissionRepositoryInterface
}

func NewGetRolePermissionsUseCase() *GetRolePermissionsUseCase {
	return &GetRolePermissionsUseCase{
		repo: repository.NewRolePermissionRepository(),
	}
}

func (uc *GetRolePermissionsUseCase) Execute(input GetRolePermissionsInputDTO) (output []GetRolePermissionsOutputDTO, err error) {
	rp, err := uc.repo.GetRolePermissionsByRole(input.RoleID)
	if err != nil {
		return
	}

	for _, r := range rp {
		output = append(output, GetRolePermissionsOutputDTO{
			ID:     r.GetRolePermissionID(),
			RoleID: r.GetRoleID(),
		})
	}
	return
}
