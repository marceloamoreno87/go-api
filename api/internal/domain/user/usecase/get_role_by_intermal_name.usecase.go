package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRoleByInternalNameInputDTO struct {
	InternalName string `json:"internal_name"`
}

type GetRoleByInternalNameOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRoleByInternalNameUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewGetRoleByInternalNameUseCase(db config.SQLCInterface) *GetRoleByInternalNameUseCase {
	return &GetRoleByInternalNameUseCase{
		repo: repository.NewRoleRepository(db),
	}
}

func (uc *GetRoleByInternalNameUseCase) Execute(ctx context.Context, input GetRoleByInternalNameInputDTO) (output GetRoleByInternalNameOutputDTO, err error) {

	role, err := uc.repo.GetRoleByInternalName(ctx, input.InternalName)
	if err != nil {
		return
	}

	output = GetRoleByInternalNameOutputDTO{
		ID:           role.GetID(),
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
		CreatedAt:    role.GetCreatedAt(),
		UpdatedAt:    role.GetUpdatedAt(),
	}

	return
}
