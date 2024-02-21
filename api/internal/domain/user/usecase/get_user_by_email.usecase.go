package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserByEmailInputDTO struct {
	Email string `json:"email"`
}

type GetUserByEmailOutputDTO struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByEmailUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewGetUserByEmailUseCase(repo repository.UserRepositoryInterface) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		repo: repo,
	}
}

func (uc *GetUserByEmailUseCase) Execute(input GetUserByEmailInputDTO) (output GetUserByEmailOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		return
	}

	output = GetUserByEmailOutputDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    user.RoleID,
		AvatarID:  user.AvatarID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
