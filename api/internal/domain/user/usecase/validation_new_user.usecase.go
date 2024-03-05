package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type ValidationNewUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type ValidationNewUserOutputDTO struct {
	Hash string `json:"hash"`
}

type ValidationNewUserUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewValidationNewUserUseCase(repo repository.UserRepositoryInterface) *ValidationNewUserUseCase {
	return &ValidationNewUserUseCase{
		repo: repo,
	}
}

func (uc *ValidationNewUserUseCase) Execute(input ValidationNewUserInputDTO) (output ValidationNewUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	user.SetActive(true)
	if err = uc.repo.CreateUser(user); err != nil {
		return
	}

	output.Hash = user.GetHash()
	return
}
