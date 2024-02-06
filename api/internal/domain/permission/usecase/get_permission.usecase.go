package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type GetPermissionInputDTO struct {
	ID int32 `json:"id"`
}

type GetPermissionOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewGetPermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *GetPermissionUseCase {
	return &GetPermissionUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *GetPermissionUseCase) Execute(input GetPermissionInputDTO) (output GetPermissionOutputDTO, err error) {
	permission, err := uc.PermissionRepository.GetPermission(input.ID)
	if err != nil {
		return GetPermissionOutputDTO{}, err
	}

	output = GetPermissionOutputDTO{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
		CreatedAt:    permission.CreatedAt,
		UpdatedAt:    permission.UpdatedAt,
	}
	return
}
