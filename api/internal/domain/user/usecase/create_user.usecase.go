package usecase

import (
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
	user.HashPassword()
	err = uc.repo.CreateUser(user)

	output = CreateUserOutputDTO{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	}
	return
}
