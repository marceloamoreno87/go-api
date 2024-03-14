package usecase

import (
	"errors"
	"log/slog"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAuthByRefreshTokenInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAuthByRefreshTokenOutputDTO struct {
	Token string `json:"token"`
}

type CreateAuthByRefreshTokenUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewCreateAuthByRefreshTokenUseCase() *CreateAuthByRefreshTokenUseCase {
	return &CreateAuthByRefreshTokenUseCase{
		repo: repository.NewUserRepository(),
	}
}

func (uc *CreateAuthByRefreshTokenUseCase) Execute(input CreateAuthByRefreshTokenInputDTO) (output CreateAuthByRefreshTokenOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		slog.Info("err", err)
		return
	}

	if !user.ComparePassword(input.Password) {
		return CreateAuthByRefreshTokenOutputDTO{}, errors.New("invalid credentials")
	}

	user.GenerateToken()

	return CreateAuthByRefreshTokenOutputDTO{Token: user.GetToken()}, nil
}
