package usecase

import (
	"errors"

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
	id, err := entity.NewAuth().ValidateToken(config.TokenAuth, input.Token)
	if err != nil {
		return GetRefreshJWTOutputDTO{}, errors.New("Not Authorized")
	}

	user, err := uc.UserRepository.GetUserByID(id)
	if err != nil {
		return GetRefreshJWTOutputDTO{}, errors.New("Not Authorized")
	}

	token, err := entity.NewAuth().NewToken(config.TokenAuth, config.Environment.JWTExpiresIn, user.GetID())
	if err != nil {
		return GetRefreshJWTOutputDTO{}, errors.New("Not Authorized")
	}

	output = GetRefreshJWTOutputDTO{
		Token: token,
	}

	return
}
