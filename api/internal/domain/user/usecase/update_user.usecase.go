package usecase

import (
	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type UpdateUserOutputDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type UpdateUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUpdateUserUseCase(userRepository repository.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (output UpdateUserOutputDTO, err error) {
	user := entity.User{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Photo:    input.Photo,
	}

	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	u, err := uc.UserRepository.UpdateUser(&user)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	output = UpdateUserOutputDTO{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Photo:    u.Photo,
	}

	return
}
