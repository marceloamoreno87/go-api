package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeletePermissionInputDTO struct {
	ID int32 `json:"id"`
}

type DeletePermissionOutputDTO struct {
	ID int32 `json:"id"`
}

type DeletePermissionUseCase struct {
	repo repository.Permissionrepository
}

func NewDeletePermissionUseCase(db config.SQLCInterface) *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		repo: repository.NewPermissionRepository(db),
	}
}

func (uc *DeletePermissionUseCase) Execute(ctx context.Context, input DeletePermissionInputDTO) (output DeletePermissionOutputDTO, err error) {
	err = uc.repo.DeletePermission(ctx, input.ID)
	output = DeletePermissionOutputDTO{
		input.ID,
	}
	return
}
