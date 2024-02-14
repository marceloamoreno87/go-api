package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewDeleteUserUseCase(userRepository repository.UserRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (err error) {
	user, err := uc.UserRepository.GetUser(input.ID)
	if err != nil {
		return
	}

	err = uc.UserRepository.DeleteUser(user.GetID())
	if err != nil {
		return
	}

	return
}
