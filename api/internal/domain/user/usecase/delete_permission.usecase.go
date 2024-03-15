package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
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
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewDeletePermissionUseCase() *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		repo: repository.NewPermissionRepository(),
	}
}

func (uc *DeletePermissionUseCase) Execute(input DeletePermissionInputDTO) (output DeletePermissionOutputDTO, err error) {
	permission, err := uc.repo.GetPermission(input.ID)
	if err != nil {
		return
	}

	p, err := uc.repo.DeletePermission(permission.GetID())
	if err != nil {
		return
	}
	output = DeletePermissionOutputDTO{
		ID:           p.GetID(),
		Name:         p.GetName(),
		InternalName: p.GetInternalName(),
		Description:  p.GetDescription(),
		CreatedAt:    p.GetCreatedAt(),
		UpdatedAt:    p.GetUpdatedAt(),
	}
	return
}
