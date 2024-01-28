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
	UserRepository repository.UserRepositoryInterface
}

func NewGetRefreshJWTUseCase(userRepository repository.UserRepositoryInterface) *GetRefreshJWTUseCase {
	return &GetRefreshJWTUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetRefreshJWTUseCase) Execute(input GetRefreshJWTInputDTO) (output GetRefreshJWTOutputDTO, err error) {
	auth := entity.NewAuth()
	if err != nil {
		return GetRefreshJWTOutputDTO{}, err
	}

	err = auth.RefreshToken(config.TokenAuth, input.Token)
	if err != nil {
		return GetRefreshJWTOutputDTO{}, err
	}

	user, err := uc.UserRepository.GetUser(auth.GetId())
	if err != nil {
		return GetRefreshJWTOutputDTO{}, err
	}

	err = auth.NewToken(config.TokenAuth, config.Environment.JWTExpiresIn, user.GetID())
	if err != nil {
		return GetRefreshJWTOutputDTO{}, err
	}

	output = GetRefreshJWTOutputDTO{
		Token: auth.GetToken(),
	}

	return
}
