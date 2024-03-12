package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/domain/auth/event"
	"github.com/marceloamoreno/goapi/internal/domain/auth/repository"
)

type ForgotPasswordInputDTO struct {
	Email string `json:"email"`
}

type ForgotPasswordUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewForgotPasswordUseCase(repo repository.UserRepositoryInterface) *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		repo: repo,
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
