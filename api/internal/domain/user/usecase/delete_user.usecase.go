package usecase

import "github.com/marceloamoreno/goapi/internal/domain/user/repository"

type DeleteUserInputDTO struct {
	ID int64 `json:"id"`
}

type DeleteUserOutputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewDeleteUserUseCase(userRepository repository.UserRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error) {
	user, err := uc.UserRepository.GetUser(input.ID)
	if err != nil {
		return DeleteUserOutputDTO{}, err
	}

	err = uc.UserRepository.DeleteUser(input.ID)
	if err != nil {
		return DeleteUserOutputDTO{}, err
	}

	output = DeleteUserOutputDTO{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return
}
