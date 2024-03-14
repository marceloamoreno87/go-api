package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdatePasswordUserInputDTO struct {
	Hash     string `json:"hash"`
	Password string `json:"password"`
}

type UpdatePasswordUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdatePasswordUserUseCase() *UpdatePasswordUserUseCase {
	return &UpdatePasswordUserUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UpdatePasswordUserUseCase) Execute(input UpdatePasswordUserInputDTO) (err error) {
	userValidation, err := uc.repo.GetValidationUserByHash(input.Hash)
	if err != nil {
		return
	}

	if !userValidation.ValidateHashExpiresIn() {
		return
	}

	user, err := uc.repo.GetUser(userValidation.GetUserID())
	if err != nil {
		return
	}
	user.SetPassword(input.Password)
	user.HashPassword()

	err = uc.repo.UpdatePasswordUser(user, user.GetID())
	if err != nil {
		return
	}

	userValidation.SetUsed(true)
	err = uc.repo.SetUserValidationUsed(userValidation.GetID())
	if err != nil {
		return
	}

	go event.NewUpdatedPasswordEmailEvent(user).Send()

	return
}
