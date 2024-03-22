package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserOutputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewDeleteUserUseCase(db config.SQLCInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error) {
	err = uc.repo.DeleteUser(ctx, input.ID)
	if err != nil {
		return
	}

	output = DeleteUserOutputDTO{
		ID: input.ID,
	}
	return
}
