package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
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
	Active    bool      `json:"active"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUsersUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewGetUsersUseCase(db config.SQLCInterface) *GetUsersUseCase {
	return &GetUsersUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *GetUsersUseCase) Execute(ctx context.Context, input GetUsersInputDTO) (output []GetUsersOutputDTO, err error) {
	users, err := uc.repo.GetUsers(ctx, input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, user := range users {
		output = append(output, GetUsersOutputDTO{
			ID:        user.GetID(),
			Email:     user.GetEmail(),
			Name:      user.GetName(),
			Password:  user.GetPassword(),
			Active:    user.GetActive(),
			RoleID:    user.GetRoleID(),
			AvatarID:  user.GetAvatarID(),
			CreatedAt: user.GetCreatedAt(),
			UpdatedAt: user.GetUpdatedAt(),
		})
	}

	return
}
