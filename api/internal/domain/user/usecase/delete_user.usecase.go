package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewDeleteUserUseCase() *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (err error) {
	user, err := uc.repo.GetUser(input.ID)
	if err != nil {
		return
	}

	if err = uc.repo.DeleteUser(user.GetID()); err != nil {
		return
	}

	return
}
