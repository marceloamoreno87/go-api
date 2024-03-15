package usecase

import (
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

func NewDeleteUserUseCase(DB config.SQLCInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repository.NewUserRepository(DB),
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error) {
	err = uc.repo.DeleteUser(input.ID)
	if err != nil {
		return
	}

	output = DeleteUserOutputDTO{
		ID: input.ID,
	}
	return
}
