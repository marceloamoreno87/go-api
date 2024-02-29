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
	repo repository.PermissionRepositoryInterface
}

func NewDeletePermissionUseCase(repo repository.PermissionRepositoryInterface) *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		repo: repo,
	}
}

func (uc *DeletePermissionUseCase) Execute(input DeletePermissionInputDTO) (err error) {
	permission, err := uc.repo.GetPermission(input.ID)
	if err != nil {
		return
	}

	if err = uc.repo.DeletePermission(permission.GetID()); err != nil {
		return
	}

	return
}
