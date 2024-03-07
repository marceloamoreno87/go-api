package usecase

import (
	"errors"

	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UserVerifyInputDTO struct {
	Hash string `json:"hash"`
}

type UserVerifyUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewUserVerifyUseCase(
	repo repository.UserRepositoryInterface,
) *UserVerifyUseCase {
	return &UserVerifyUseCase{
		repo: repo,
	}
}

func (uc *UserVerifyUseCase) Execute(input UserVerifyInputDTO) (err error) {
	userValidation, err := uc.repo.GetValidationUserByHash(input.Hash)
	if err != nil {
		return
	}

	if userValidation == nil {
		return
	}

	if !userValidation.ValidateHashExpiresIn() {
		return errors.New("hash expired")
	}

	user, err := uc.repo.GetUser(userValidation.GetUserID())
	if err != nil {
		return
	}

	user.SetActive(true)
	err = uc.repo.UpdateUser(user, user.ID)
	if err != nil {
		return
	}

	userValidation.SetUsed(true)
	err = uc.repo.UpdateValidationUser(userValidation, userValidation.ID)
	if err != nil {
		return
	}

	userValidation.SetUser(user)

	go event.NewUserVerifiedEmailEvent(userValidation).Send()

	return
}
