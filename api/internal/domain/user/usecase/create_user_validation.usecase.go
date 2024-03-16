package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserValidationInputDTO struct {
	UserID int32 `json:"user_id"`
}

type CreateUserValidationOutputDTO struct {
	UserID    int32  `json:"user_id"`
	Hash      string `json:"hash"`
	ExpiresIn int32  `json:"expires_in"`
	Used      bool   `json:"used"`
}

type CreateUserValidationUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewCreateUserValidationUseCase() *CreateUserValidationUseCase {
	return &CreateUserValidationUseCase{
		repo: repository.NewUserValidationRepository(),
	}
}

func (uc *CreateUserValidationUseCase) Execute(input CreateUserValidationInputDTO) (output CreateUserValidationOutputDTO, err error) {
	userValidation, err := entity.NewUserValidation(input.UserID)
	if err != nil {
		return
	}

	err = uc.repo.CreateValidationUser(userValidation)
	if err != nil {
		return
	}

	output = CreateUserValidationOutputDTO{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
		Used:      userValidation.GetUsed(),
	}

	return
}
