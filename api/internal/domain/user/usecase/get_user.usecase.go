package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
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
	repo repositoryInterface.UserRepositoryInterface
}

func NewGetUserUseCase() *GetUserUseCase {
	return &GetUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *GetUserUseCase) Execute(input GetUserInputDTO) (output GetUserOutputDTO, err error) {
	user, err := uc.repo.GetUser(input.ID)
	if err != nil {
		return
	}

	output = GetUserOutputDTO{
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
