package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type GetPermissionsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPermissionsOutputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type GetPermissionsUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewGetPermissionsUseCase(permissionRepository repository.PermissionRepositoryInterface) *GetPermissionsUseCase {
	return &GetPermissionsUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *GetPermissionsUseCase) Execute(input GetPermissionsInputDTO) (output []GetPermissionsOutputDTO, err error) {

	permissions, err := uc.PermissionRepository.GetPermissions(input.Limit, input.Offset)
	if err != nil {
		return []GetPermissionsOutputDTO{}, err
	}

	for _, permission := range permissions {
		output = append(output, GetPermissionsOutputDTO{
			ID:           permission.ID,
			Name:         permission.Name,
			InternalName: permission.InternalName,
			Description:  permission.Description,
		})
	}

	return
}
