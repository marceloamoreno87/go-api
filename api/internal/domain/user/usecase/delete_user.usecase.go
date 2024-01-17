package usecase

import "github.com/marceloamoreno/izimoney/internal/domain/user/repository"

type DeleteUserInputDTO struct {
	ID int64 `json:"id"`
}

type DeleteUserOutputDTO struct {
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
	_, err = uc.UserRepository.GetUser(input.ID)
	if err != nil {
		return
	}

	err = uc.UserRepository.DeleteUser(input.ID)
	if err != nil {
		return
	}

	return
}
