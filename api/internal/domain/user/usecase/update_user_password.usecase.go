package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserPasswordInputDTO struct {
}

type UpdateUserPasswordUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserPasswordUseCase() *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) (err error) {

	// err = uc.repo.UpdateUserPassword(user, user.GetID())
	// if err != nil {
	// 	return
	// }

	return
}
