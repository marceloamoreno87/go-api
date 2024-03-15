package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserOutputDTO struct {
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

type DeleteUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewDeleteUserUseCase() *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error) {
	user, err := uc.repo.GetUser(input.ID)
	if err != nil {
		return
	}

	u, err := uc.repo.DeleteUser(user.GetID())
	if err != nil {
		return
	}

	output = DeleteUserOutputDTO{
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
