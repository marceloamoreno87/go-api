package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleOutputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleUseCase struct {
	repo repository.Rolerepository
}

func NewDeleteRoleUseCase(db config.SQLCInterface) *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		repo: repository.NewRoleRepository(db),
	}
}

func (uc *DeleteRoleUseCase) Execute(ctx context.Context, input DeleteRoleInputDTO) (output DeleteRoleOutputDTO, err error) {
	err = uc.repo.DeleteRole(ctx, input.ID)
	if err != nil {
		return
	}
	output = DeleteRoleOutputDTO{
		ID: input.ID,
	}
	return
}
