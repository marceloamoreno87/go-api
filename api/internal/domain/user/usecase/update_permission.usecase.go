package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
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
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewUpdatePermissionUseCase() *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		repo: repository.NewPermissionRepository(),
	}
}

func (uc *UpdatePermissionUseCase) Execute(input UpdatePermissionInputDTO) (output UpdatePermissionOutputDTO, err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	p, err := uc.repo.UpdatePermission(permission, input.ID)
	if err != nil {
		return
	}
	output = UpdatePermissionOutputDTO{
		ID:           p.GetID(),
		Name:         p.GetName(),
		InternalName: p.GetInternalName(),
		Description:  p.GetDescription(),
		CreatedAt:    p.GetCreatedAt(),
		UpdatedAt:    p.GetUpdatedAt(),
	}
	return
}
