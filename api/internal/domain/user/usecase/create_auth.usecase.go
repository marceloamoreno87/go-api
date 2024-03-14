package usecase

import (
	"errors"
	"log/slog"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAuthInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAuthOutputDTO struct {
	Token string `json:"token"`
}

type CreateAuthUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewAuthUseCase() *CreateAuthUseCase {
	return &CreateAuthUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *CreateAuthUseCase) Execute(input CreateAuthInputDTO) (output CreateAuthOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		slog.Info("err", err)
		return
	}

	if !user.ComparePassword(input.Password) {
		return CreateAuthOutputDTO{}, errors.New("invalid credentials")
	}

	user.GenerateToken()

	return CreateAuthOutputDTO{Token: user.GetToken()}, nil
}
