package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type ForgotPasswordInputDTO struct {
	Email string `json:"email"`
}

type ForgotPasswordUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewForgotPasswordUseCase() *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *ForgotPasswordUseCase) Execute(input ForgotPasswordInputDTO) (err error) {

	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		return
	}

	userValidation, err := entity.NewUserValidation(user)
	if err != nil {
		return
	}

	err = uc.repo.CreateValidationUser(userValidation)
	if err != nil {
		return
	}

	go event.NewPasswordForgotEmailEvent(userValidation).Send()

	return
}
