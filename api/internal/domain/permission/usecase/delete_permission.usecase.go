package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type DeletePermissionInputDTO struct {
	ID int32 `json:"id"`
}

type DeletePermissionOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeletePermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewDeletePermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		PermissionRepository: permissionRepository,
	}
}

func (uc *DeletePermissionUseCase) Execute(input DeletePermissionInputDTO) (output DeletePermissionOutputDTO, err error) {
	permission, err := uc.PermissionRepository.GetPermission(input.ID)
	if err != nil {
		return DeletePermissionOutputDTO{}, err
	}

	u, err := uc.PermissionRepository.DeletePermission(permission.GetID())
	if err != nil {
		return DeletePermissionOutputDTO{}, err
	}

	output = DeletePermissionOutputDTO{
		ID:           u.ID,
		Name:         u.Name,
		InternalName: u.InternalName,
		Description:  u.Description,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
	return
}
