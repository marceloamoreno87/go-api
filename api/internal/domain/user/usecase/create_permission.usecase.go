package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
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
	repo repository.Permissionrepository
}

func NewCreatePermissionUseCase(db config.SQLCInterface) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		repo: repository.NewPermissionRepository(db),
	}
}

func (uc *CreatePermissionUseCase) Execute(ctx context.Context, input CreatePermissionInputDTO) (output CreatePermissionOutputDTO, err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.repo.CreatePermission(ctx, permission)

	output = CreatePermissionOutputDTO{
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	}
	return
}
