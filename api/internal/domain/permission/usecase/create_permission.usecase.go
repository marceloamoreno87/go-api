package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type CreatePermissionInputDTO struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type CreatePermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewCreatePermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *CreatePermissionUseCase) Execute(input CreatePermissionInputDTO) (err error) {
	permission := &entity.Permission{
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	}

	err = uc.PermissionRepository.CreatePermission(permission)
	if err != nil {
		return
	}

	return
}
