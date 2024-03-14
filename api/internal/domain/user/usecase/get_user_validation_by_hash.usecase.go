package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserValidationByHashInputDTO struct {
}

type GetUserValidationByHashOutputDTO struct {
}

type GetUserValidationByHashUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewGetUserValidationByHashUseCase() *GetUserValidationByHashUseCase {
	return &GetUserValidationByHashUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *GetUserValidationByHashUseCase) Execute(input GetUserInputDTO) (output GetUserOutputDTO, err error) {
	// userValidation, err := uc.repo.GetValidationUserByHash(input.Hash)
	// if err != nil {
	// 	return
	// }

	// if !userValidation.ValidateHashExpiresIn() {
	// 	return
	// }

	return
}
