package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type UpdatePermissionInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdatePermissionUseCase struct {
	repo repository.PermissionRepositoryInterface
}

func NewUpdatePermissionUseCase(repo repository.PermissionRepositoryInterface) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		repo: repo,
	}
}

func (uc *UpdatePermissionUseCase) Execute(input UpdatePermissionInputDTO) (err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	if err = uc.repo.UpdatePermission(permission, input.ID); err != nil {
		return
	}

	return
}
