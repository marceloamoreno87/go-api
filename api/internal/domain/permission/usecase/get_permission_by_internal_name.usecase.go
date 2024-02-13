package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type GetPermissionByInternalNameInputDTO struct {
	InternalName string `json:"internal_name"`
}

type GetPermissionByInternalNameOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionByInternalNameUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewGetPermissionByInternalNameUseCase(permissionRepository repository.PermissionRepositoryInterface) *GetPermissionByInternalNameUseCase {
	return &GetPermissionByInternalNameUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *GetPermissionByInternalNameUseCase) Execute(input GetPermissionByInternalNameInputDTO) (output GetPermissionByInternalNameOutputDTO, err error) {
	permission, err := uc.PermissionRepository.GetPermissionByInternalName(input.InternalName)
	if err != nil {
		return
	}

	output = GetPermissionByInternalNameOutputDTO{
		ID:           permission.ID,
		Name:         permission.Name,
		InternalName: permission.InternalName,
		Description:  permission.Description,
		CreatedAt:    permission.CreatedAt,
		UpdatedAt:    permission.UpdatedAt,
	}
	return
}
