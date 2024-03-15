package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserPasswordInputDTO struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

type UpdateUserPasswordOutputDTO struct {
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

type UpdateUserPasswordUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserPasswordUseCase() *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) (output UpdateUserPasswordOutputDTO, err error) {
	u, err := uc.repo.UpdateUserPassword(input.ID, input.Password)
	if err != nil {
		return
	}

	output = UpdateUserPasswordOutputDTO{
		ID:        u.GetID(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Password:  u.GetPassword(),
		Active:    u.GetActive(),
		RoleID:    u.GetRoleID(),
		AvatarID:  u.GetAvatarID(),
		CreatedAt: u.GetCreatedAt(),
		UpdatedAt: u.GetUpdatedAt(),
	}

	return
}
