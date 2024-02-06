package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type CreatePermissionInputDTO struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type CreatePermissionOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreatePermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewCreatePermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *CreatePermissionUseCase) Execute(input CreatePermissionInputDTO) (output CreatePermissionOutputDTO, err error) {
	permission := &entity.Permission{
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	}

	permission, err = uc.PermissionRepository.CreatePermission(permission)
	if err != nil {
		return CreatePermissionOutputDTO{}, err
	}

	output = CreatePermissionOutputDTO{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
		CreatedAt:    permission.CreatedAt,
		UpdatedAt:    permission.UpdatedAt,
	}

	return
}
