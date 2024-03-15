package usecase

import (
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
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
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

	err = uc.repo.CreatePermission(permission)

	output = CreatePermissionOutputDTO{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	}
	return
}
