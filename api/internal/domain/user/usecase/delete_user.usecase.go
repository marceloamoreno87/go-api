package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int64 `json:"id"`
}

type DeleteUserOutputDTO struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return
}
