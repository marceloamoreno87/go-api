package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetJWTInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutputDTO struct {
	Token string `json:"token"`
}

type GetJWTUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewGetJWTUseCase(userRepository repository.UserRepositoryInterface) *GetJWTUseCase {
	return &GetJWTUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetJWTUseCase) Execute(input GetJWTInputDTO) (output GetJWTOutputDTO, err error) {

	auth := entity.NewAuth()
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	user, err := uc.UserRepository.GetUserByEmail(input.Email)
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	if !user.ComparePassword(input.Password) {
		return GetJWTOutputDTO{}, err
	}

	err = auth.NewToken(config.TokenAuth, config.Environment.JWTExpiresIn, user.GetID())
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	output = GetJWTOutputDTO{
		Token: auth.GetToken(),
	}

	return
}
