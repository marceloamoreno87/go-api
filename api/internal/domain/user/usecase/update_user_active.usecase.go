package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserActiveInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UpdateUserActiveOutputDTO struct {
}

type UpdateUserActiveUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserActiveUseCase() *UpdateUserActiveUseCase {
	return &UpdateUserActiveUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdateUserActiveUseCase) Execute(input UpdateUserActiveInputDTO) (err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}
	user.SetID(input.ID)
	user.SetActive(input.Active)

	err = uc.repo.UpdateUserActive(user.GetID(), user.GetActive())

	return
}
