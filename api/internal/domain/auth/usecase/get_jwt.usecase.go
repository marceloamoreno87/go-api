package usecase

import (
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
}

// TODO: REFACTOR
func NewGetJWTUseCase(repo repository.UserRepositoryInterface) *GetJWTUseCase {
	return &GetJWTUseCase{
		GetUserByEmailUseCase: usecase.NewGetUserByEmailUseCase(repo),
	}
}

func (uc *GetJWTUseCase) Execute(input GetJWTInputDTO) (output GetJWTOutputDTO, err error) {

	auth := AuthEntity.NewAuth()
	if err != nil {
		return
	}
	inputGetUser := usecase.GetUserByEmailInputDTO{Email: input.Email}
	user, err := uc.GetUserByEmailUseCase.Execute(inputGetUser)
	if err != nil {
		return
	}

	newUser := &UserEntity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if !newUser.ComparePassword(input.Password) {
		return
	}

	err = auth.NewToken(config.TokenAuth, config.Environment.JWTExpiresIn, newUser.GetID())
	if err != nil {
		return
	}

	output.Token = auth.GetToken()
	return
}
