package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRefreshJWTInputDTO struct {
	Token string `json:"token"`
}

type GetRefreshJWTOutputDTO struct {
	Token string `json:"token"`
}

type GetRefreshJWTUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewGetRefreshJWTUseCase(repo repository.UserRepositoryInterface) *GetRefreshJWTUseCase {
	return &GetRefreshJWTUseCase{
		repo: repo,
	}
}

func (uc *GetRefreshJWTUseCase) Execute(input GetRefreshJWTInputDTO) (output GetRefreshJWTOutputDTO, err error) {
	auth := entity.NewAuth()
	if err != nil {
		return
	}

	err = auth.RefreshToken(config.NewTokenAuth(), input.Token)
	if err != nil {
		return
	}

	user, err := uc.repo.GetUser(auth.GetID())
	if err != nil {
		return
	}

	err = auth.NewToken(config.NewTokenAuth(), config.NewTokenAuth().GetJWTExpiresIn(), user.GetID())
	if err != nil {
		return
	}

	output = GetRefreshJWTOutputDTO{
		Token: auth.GetToken(),
	}

	return
}
