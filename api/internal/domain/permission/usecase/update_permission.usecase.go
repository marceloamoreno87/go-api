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
	ID   int32
}

func NewUpdatePermissionUseCase(repo repository.PermissionRepositoryInterface, id int32) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		repo: repo,
		ID:   id,
	}
}

func (uc *UpdatePermissionUseCase) Execute(input UpdatePermissionInputDTO) (err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.repo.UpdatePermission(permission, uc.ID)
	if err != nil {
		return
	}

	return
}
