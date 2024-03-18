package usecase

import (
	"errors"

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
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
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
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}

	check, _ := uc.repo.GetUserByEmail(user.GetEmail())
	if check != nil && check.GetID() != 0 {
		return output, errors.New("user already exists")
	}

	user.HashPassword()
	newUser, err := uc.repo.CreateUser(user)
	if err != nil {
		return
	}

	output = CreateUserOutputDTO{
		ID:       newUser.GetID(),
		Name:     newUser.GetName(),
		Email:    newUser.GetEmail(),
		Password: newUser.GetPassword(),
		Active:   newUser.GetActive(),
		RoleID:   newUser.GetRoleID(),
		AvatarID: newUser.GetAvatarID(),
	}
	return
}
