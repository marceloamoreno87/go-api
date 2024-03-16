package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserPasswordInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordOutputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}
	user.SetID(input.ID)
	user.HashPassword()

	err = uc.repo.UpdateUserPassword(user.GetID(), user.GetPassword())
	output = UpdateUserPasswordOutputDTO{
		ID:       input.ID,
		Password: input.Password,
	}

	return
}
