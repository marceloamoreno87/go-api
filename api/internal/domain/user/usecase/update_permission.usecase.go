package usecase

import (
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

type UpdatePermissionUseCase struct {
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewUpdatePermissionUseCase() *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		repo: repository.NewPermissionRepository(),
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
