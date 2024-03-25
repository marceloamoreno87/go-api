package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRolesInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetRolesOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRolesUseCase struct {
	repo repository.Rolerepository
}

func NewGetRolesUseCase(db config.SQLCInterface) *GetRolesUseCase {
	return &GetRolesUseCase{
		repo: repository.NewRoleRepository(db),
	}
}

func (uc *GetRolesUseCase) Execute(ctx context.Context, input GetRolesInputDTO) (output []GetRolesOutputDTO, err error) {

	roles, err := uc.repo.GetRoles(ctx, input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, role := range roles {
		output = append(output, GetRolesOutputDTO{
			ID:           role.GetID(),
			Name:         role.GetName(),
			InternalName: role.GetInternalName(),
			Description:  role.GetDescription(),
			CreatedAt:    role.GetCreatedAt(),
			UpdatedAt:    role.GetUpdatedAt(),
		})
	}

	return
}
