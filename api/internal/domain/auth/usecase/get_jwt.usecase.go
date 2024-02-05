package usecase

import (
	"errors"

	"github.com/marceloamoreno/goapi/config"
	AuthEntity "github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	UserEntity "github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetJWTInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutputDTO struct {
	Token string `json:"token"`
}

type GetJWTUseCase struct {
	GetUserByEmailUseCase usecase.GetUserByEmailUseCaseInterface
	UserRepository        repository.UserRepositoryInterface
}

func NewGetJWTUseCase(UserRepository repository.UserRepositoryInterface) *GetJWTUseCase {
	return &GetJWTUseCase{
		GetUserByEmailUseCase: usecase.NewGetUserByEmailUseCase(UserRepository),
	}
}

func (uc *GetJWTUseCase) Execute(input GetJWTInputDTO) (output GetJWTOutputDTO, err error) {

	auth := AuthEntity.NewAuth()
	if err != nil {
		return GetJWTOutputDTO{}, err
	}

	user, err := uc.GetUserByEmailUseCase.Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if err != nil {
		return GetJWTOutputDTO{}, errors.New("not authorized")
	}

	newUser := &UserEntity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleId:    user.RoleId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if !newUser.ComparePassword(input.Password) {
		return GetJWTOutputDTO{}, errors.New("not authorized")
	}

	err = auth.NewToken(config.TokenAuth, config.Environment.JWTExpiresIn, newUser.GetID())
	if err != nil {
		return GetJWTOutputDTO{}, errors.New("not authorized")
	}

	output = GetJWTOutputDTO{
		Token: auth.GetToken(),
	}

	return
}
