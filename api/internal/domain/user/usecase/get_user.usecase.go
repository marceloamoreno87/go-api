package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserInputDTO struct {
	ID int32 `json:"id"`
}

type GetUserOutputDTO struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewGetUserUseCase(repo repository.UserRepositoryInterface) *GetUserUseCase {
	return &GetUserUseCase{
		repo: repo,
	}
}

func (uc *GetUserUseCase) Execute(input GetUserInputDTO) (output GetUserOutputDTO, err error) {
	user, err := uc.repo.GetUser(input.ID)
	if err != nil {
		return
	}

	output = GetUserOutputDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Active:    user.Active,
		RoleID:    user.RoleID,
		AvatarID:  user.AvatarID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
