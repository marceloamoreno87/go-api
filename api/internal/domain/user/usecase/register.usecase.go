package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type RegisterInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type RegisterOutputDTO struct {
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

type RegisterUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewRegisterUseCase() *RegisterUseCase {
	return &RegisterUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *RegisterUseCase) Execute(input RegisterInputDTO) (output RegisterOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, 1, 1)
	if err != nil {
		return
	}

	newUser, err := uc.repo.RegisterUser(user)
	if err != nil {
		return
	}

	output = RegisterOutputDTO{
		ID:        newUser.GetID(),
		Name:      newUser.GetName(),
		Email:     newUser.GetEmail(),
		Password:  newUser.GetPassword(),
		Active:    newUser.GetActive(),
		RoleID:    newUser.GetRoleID(),
		AvatarID:  newUser.GetAvatarID(),
		CreatedAt: newUser.GetCreatedAt(),
		UpdatedAt: newUser.GetUpdatedAt(),
	}

	return
}
