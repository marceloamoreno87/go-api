package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUsersInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersOutputDTO struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	RoleID    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUsersUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewGetUsersUseCase(userRepository repository.UserRepositoryInterface) *GetUsersUseCase {
	return &GetUsersUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUsersUseCase) Execute(input GetUsersInputDTO) (output []GetUsersOutputDTO, err error) {
	users, err := uc.UserRepository.GetUsers(input.Limit, input.Offset)
	if err != nil {
		return []GetUsersOutputDTO{}, err
	}

	for _, user := range users {
		output = append(output, GetUsersOutputDTO{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Password:  user.Password,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return
}
