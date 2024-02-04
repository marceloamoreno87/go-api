package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserByEmailInputDTO struct {
	Email string `json:"email"`
}

type GetUserByEmailOutputDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByEmailUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewGetUserByEmailUseCase(userRepository repository.UserRepositoryInterface) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUserByEmailUseCase) Execute(input GetUserByEmailInputDTO) (output GetUserByEmailOutputDTO, err error) {
	user, err := uc.UserRepository.GetUserByEmail(input.Email)
	if err != nil {
		return GetUserByEmailOutputDTO{}, err
	}

	output = GetUserByEmailOutputDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
