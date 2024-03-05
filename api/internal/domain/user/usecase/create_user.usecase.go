package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
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
	repo repository.UserRepositoryInterface
}

func NewCreateUserUseCase(repo repository.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
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
