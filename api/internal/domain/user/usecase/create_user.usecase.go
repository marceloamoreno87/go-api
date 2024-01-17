package usecase

import (
	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
)

type CreateUserInputDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type CreateUserOutputDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type CreateUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository repository.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (output CreateUserOutputDTO, err error) {

	user, err := entity.NewUser(input.Username, input.Password, input.Photo)

	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	u, err := uc.UserRepository.CreateUser(user)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	output = CreateUserOutputDTO{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Photo:    u.Photo,
	}

	return
}
