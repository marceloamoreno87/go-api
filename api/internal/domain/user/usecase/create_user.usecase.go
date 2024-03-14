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

type CreateUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewCreateUserUseCase() *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	user.SetActive(true)
	if err = uc.repo.CreateUser(user); err != nil {
		return
	}

	return
}
