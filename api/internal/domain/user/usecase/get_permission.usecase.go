package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetPermissionInputDTO struct {
	ID int32 `json:"id"`
}

type GetPermissionOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionUseCase struct {
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewGetPermissionUseCase(db config.SQLCInterface) *GetPermissionUseCase {
	return &GetPermissionUseCase{
		repo: repository.NewPermissionRepository(db),
	}
}

func (uc *GetPermissionUseCase) Execute(ctx context.Context, input GetPermissionInputDTO) (output GetPermissionOutputDTO, err error) {
	permission, err := uc.repo.GetPermission(ctx, input.ID)
	if err != nil {
		return
	}

	output = GetPermissionOutputDTO{
		ID:           permission.GetID(),
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
		CreatedAt:    permission.GetCreatedAt(),
		UpdatedAt:    permission.GetUpdatedAt(),
	}
	return
}
