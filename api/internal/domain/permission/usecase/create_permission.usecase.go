package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
)

type CreatePermissionInputDTO struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type CreatePermissionUseCase struct {
	repo repository.PermissionRepositoryInterface
}

func NewCreatePermissionUseCase(repo repository.PermissionRepositoryInterface) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		repo: repo,
	}
}

func (uc *CreatePermissionUseCase) Execute(input CreatePermissionInputDTO) (err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.repo.CreatePermission(permission)
	if err != nil {
		return
	}

	return
}
