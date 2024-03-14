package usecase

import (
	"errors"
	"log/slog"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	Token string `json:"token"`
}

type LoginUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *LoginUseCase) Execute(input LoginInputDTO) (output LoginOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		slog.Info("err", err)
		return
	}

	if !user.ComparePassword(input.Password) {
		return LoginOutputDTO{}, errors.New("invalid credentials")
	}

	user.GenerateToken()

	return LoginOutputDTO{Token: user.GetToken()}, nil
}
