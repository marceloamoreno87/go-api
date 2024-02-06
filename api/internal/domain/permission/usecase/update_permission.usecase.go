package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type UpdatePermissionInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdatePermissionOutputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdatePermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewUpdatePermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *UpdatePermissionUseCase) Execute(input *UpdatePermissionInputDTO) (*UpdatePermissionOutputDTO, error) {
	permission := &entity.Permission{
		ID:           input.ID,
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	}

	permission, err := uc.PermissionRepository.UpdatePermission(permission, input.ID)
	if err != nil {
		return &UpdatePermissionOutputDTO{}, err
	}

	return &UpdatePermissionOutputDTO{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
	}, nil
}
