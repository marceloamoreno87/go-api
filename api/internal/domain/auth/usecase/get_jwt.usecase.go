package usecase

import (
	"errors"
	"strconv"
	"time"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
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
		return GetJWTOutputDTO{}, errors.New("Not Authorized")
	}

	if !user.ComparePassword(input.Password) {
		return GetJWTOutputDTO{}, errors.New("Not Authorized")
	}

	jwtExpiresInStr := config.Environment.JWTExpiresIn
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	if err != nil {
		return GetJWTOutputDTO{}, errors.New("Not Authorized")
	}

	_, tokenString, err := config.TokenAuth.Encode(map[string]interface{}{
		"id":  user.GetID(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	if err != nil {
		return GetJWTOutputDTO{}, errors.New("Not Authorized")
	}

	output = GetJWTOutputDTO{
		Token: tokenString,
	}

	return
}
