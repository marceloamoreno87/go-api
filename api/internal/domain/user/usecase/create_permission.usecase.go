package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
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
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewCreatePermissionUseCase() *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		repo: repository.NewPermissionRepository(),
	}
}

func (uc *CreatePermissionUseCase) Execute(input CreatePermissionInputDTO) (output CreatePermissionOutputDTO, err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	p, err := uc.repo.CreatePermission(permission)
	if err != nil {
		return
	}

	output = CreatePermissionOutputDTO{
		ID:           p.GetID(),
		Name:         p.GetName(),
		InternalName: p.GetInternalName(),
		Description:  p.GetDescription(),
		CreatedAt:    p.GetCreatedAt(),
		UpdatedAt:    p.GetUpdatedAt(),
	}
	return
}
