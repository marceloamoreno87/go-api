package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type UpdateUserOutputDTO struct {
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

type UpdateUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserUseCase() *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (output UpdateUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	u, err := uc.repo.UpdateUser(user, input.ID)
	if err != nil {
		return
	}

	output = UpdateUserOutputDTO{
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
