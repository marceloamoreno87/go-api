package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}
type CreateUserOutputDTO struct {
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

type CreateUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewCreateUserUseCase() *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (output CreateUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	user.SetActive(true)
	u, err := uc.repo.CreateUser(user)
	if err != nil {
		return
	}
	output = CreateUserOutputDTO{
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
