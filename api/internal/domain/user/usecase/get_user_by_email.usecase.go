package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
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
	Active    bool      `json:"active"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByEmailUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewGetUserByEmailUseCase(db config.SQLCInterface) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *GetUserByEmailUseCase) Execute(ctx context.Context, input GetUserByEmailInputDTO) (output GetUserByEmailOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return
	}
	output = GetUserByEmailOutputDTO{
		ID:        user.GetID(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Password:  user.GetPassword(),
		Active:    user.GetActive(),
		RoleID:    user.GetRoleID(),
		AvatarID:  user.GetAvatarID(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}

	return
}
