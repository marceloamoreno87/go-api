package usecase

import (
	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserOutputDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
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
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	return
}
