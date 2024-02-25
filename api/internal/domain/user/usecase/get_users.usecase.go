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
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUsersUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewGetUsersUseCase(repo repository.UserRepositoryInterface) *GetUsersUseCase {
	return &GetUsersUseCase{
		repo: repo,
	}
}

func (uc *GetUsersUseCase) Execute(input GetUsersInputDTO) (output []GetUsersOutputDTO, err error) {
	users, err := uc.repo.GetUsers(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, user := range users {
		output = append(output, GetUsersOutputDTO{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Password:  user.Password,
			RoleID:    user.RoleID,
			AvatarID:  user.AvatarID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return
}
