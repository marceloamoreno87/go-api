package usecase

import (
	"time"

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
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdatePermissionUseCase struct {
	PermissionRepository repository.PermissionRepositoryInterface
	ID                   int32
}

func NewUpdatePermissionUseCase(permissionRepository repository.PermissionRepositoryInterface, id int32) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		PermissionRepository: permissionRepository,
		ID:                   id,
	}
}

func (uc *UpdatePermissionUseCase) Execute(input UpdatePermissionInputDTO) (output UpdatePermissionOutputDTO, err error) {

	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return UpdatePermissionOutputDTO{}, err
	}

	permission.SetID(input.ID)

	u, err := uc.PermissionRepository.UpdatePermission(permission, uc.ID)
	if err != nil {
		return UpdatePermissionOutputDTO{}, err
	}

	output = UpdatePermissionOutputDTO{
		ID:           u.ID,
		Name:         u.Name,
		InternalName: u.InternalName,
		Description:  u.Description,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}

	return
}
