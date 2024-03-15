package usecase

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type CreateAuthInputDTO struct {
	UserID       int32  `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateAuthOutputDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateAuthUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewAuthUseCase() *CreateAuthUseCase {
	return &CreateAuthUseCase{
		repo: repository.NewAuthRepository(),
	}
}

func (uc *CreateAuthUseCase) Execute(input CreateAuthInputDTO) (output CreateAuthOutputDTO, err error) {
	auth, err := entity.NewAuth(input.UserID, input.Token, input.RefreshToken)
	if err != nil {
		slog.Info("err", err)
		return
	}
	auth.SetActive(true)
	auth.SetExpiresIn(helper.StrToInt32(config.Environment.GetJWTExpiresIn()))

	newAuth, err := uc.repo.CreateAuth(auth)
	if err != nil {
		slog.Info("err", err)
		return
	}

	output.Token = newAuth.GetToken()
	output.RefreshToken = newAuth.GetRefreshToken()
	return
}
