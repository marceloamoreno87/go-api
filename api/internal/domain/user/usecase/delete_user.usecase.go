package usecase

import "github.com/marceloamoreno/goapi/internal/domain/user/repository"

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewDeleteUserUseCase(repo repository.UserRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repo,
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (err error) {
	user, err := uc.repo.GetUser(input.ID)
	if err != nil {
		return
	}

	err = uc.repo.DeleteUser(user.GetID())
	if err != nil {
		return
	}

	return
}
