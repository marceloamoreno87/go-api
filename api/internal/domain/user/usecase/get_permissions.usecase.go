package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetPermissionsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPermissionsOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionsUseCase struct {
	repo repository.Permissionrepository
}

func NewGetPermissionsUseCase(db config.SQLCInterface) *GetPermissionsUseCase {
	return &GetPermissionsUseCase{
		repo: repository.NewPermissionRepository(db),
	}
}

func (uc *GetPermissionsUseCase) Execute(ctx context.Context, input GetPermissionsInputDTO) (output []GetPermissionsOutputDTO, err error) {
	permissions, err := uc.repo.GetPermissions(ctx, input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, permission := range permissions {
		output = append(output, GetPermissionsOutputDTO{
			ID:           permission.GetID(),
			Name:         permission.GetName(),
			InternalName: permission.GetInternalName(),
			Description:  permission.GetDescription(),
			CreatedAt:    permission.GetCreatedAt(),
			UpdatedAt:    permission.GetUpdatedAt(),
		})
	}

	return
}
