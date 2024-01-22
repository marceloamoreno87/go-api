package usecase

import (
	"strconv"
	"time"

	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
)

type GetJWTInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutputDTO struct {
	Token interface{} `json:"token"`
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

	user, err := uc.UserRepository.GetUserByEmail(input.Email)
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	if !user.ComparePassword(input.Password) {
		return GetJWTOutputDTO{}, err
	}

	jwtExpiresInStr := config.Environment.JWTExpiresIn
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	_, tokenString, err := config.TokenAuth.Encode(map[string]interface{}{
		"id":  user.GetID(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	output = GetJWTOutputDTO{
		Token: tokenString,
	}
	return
}
