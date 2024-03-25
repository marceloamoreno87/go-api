package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdatePermissionInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdatePermissionOutputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdatePermissionUseCase struct {
	repo repository.Permissionrepository
}

func NewUpdatePermissionUseCase(db config.SQLCInterface) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		repo: repository.NewPermissionRepository(db),
	}
}

func (uc *UpdatePermissionUseCase) Execute(ctx context.Context, input UpdatePermissionInputDTO) (output UpdatePermissionOutputDTO, err error) {
	permission, err := entity.NewPermission(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.repo.UpdatePermission(ctx, permission, input.ID)
	output = UpdatePermissionOutputDTO{
		ID:           permission.GetID(),
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
	}
	return
}
