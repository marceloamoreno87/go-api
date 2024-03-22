package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserValidationInputDTO struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
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

func NewCreateUserValidationUseCase(db config.SQLCInterface) *CreateUserValidationUseCase {
	return &CreateUserValidationUseCase{
		repo: repository.NewUserValidationRepository(db),
	}
}

func (uc *CreateUserValidationUseCase) Execute(ctx context.Context, input CreateUserValidationInputDTO) (output CreateUserValidationOutputDTO, err error) {
	userValidation, err := entity.NewUserValidation(input.UserID)
	if err != nil {
		return
	}

	newUserValidation, err := uc.repo.CreateUserValidation(ctx, userValidation)
	if err != nil {
		return
	}

	output = CreateUserValidationOutputDTO{
		UserID:    newUserValidation.GetUserID(),
		Hash:      newUserValidation.GetHash(),
		ExpiresIn: newUserValidation.GetExpiresIn(),
		Used:      newUserValidation.GetUsed(),
	}

	return
}
