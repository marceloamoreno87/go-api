package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserValidationUsedInputDTO struct {
}

type UpdateUserValidationUsedUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserValidationUsedUseCase() *UpdateUserValidationUsedUseCase {
	return &UpdateUserValidationUsedUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdateUserValidationUsedUseCase) Execute(input UpdateUserValidationUsedInputDTO) (err error) {
	// err = uc.repo.UpdatedUserValidationUsed(userValidation.GetID())
	// if err != nil {
	// 	return
	// }
	return
}
