package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAuthInputDTO struct {
	UserID int32 `json:"user_id"`
}

type CreateAuthOutputDTO struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}

type CreateAuthUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewCreateAuthUseCase(DB config.SQLCInterface) *CreateAuthUseCase {
	return &CreateAuthUseCase{
		repo: repository.NewAuthRepository(DB),
	}
}

func (uc *CreateAuthUseCase) Execute(input CreateAuthInputDTO) (output CreateAuthOutputDTO, err error) {
	auth, err := entity.NewAuth(input.UserID)
	if err != nil {
		return
	}

	err = uc.repo.CreateAuth(auth)
	if err != nil {
		return
	}

	output = CreateAuthOutputDTO{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		Active:                auth.GetActive(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	}

	return
}
