package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserOutputDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
	ID             int64
}

func NewUpdateUserUseCase(userRepository repository.UserRepositoryInterface, id int64) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
		ID:             id,
	}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (output UpdateUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	u, err := uc.UserRepository.UpdateUser(user, uc.ID)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	output = UpdateUserOutputDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return
}
