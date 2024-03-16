package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetAuthByRefreshTokenInputDTO struct {
	UserID       int32  `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type GetAuthByRefreshTokenOutputDTO struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}

type GetAuthByRefreshTokenUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewGetAuthByRefreshTokenUseCase() *GetAuthByRefreshTokenUseCase {
	return &GetAuthByRefreshTokenUseCase{
		repo: repository.NewAuthRepository(),
	}
}

func (uc *GetAuthByRefreshTokenUseCase) Execute(input GetAuthByRefreshTokenInputDTO) (output GetAuthByRefreshTokenOutputDTO, err error) {
	auth, err := uc.repo.GetAuthByRefreshToken(input.UserID, input.RefreshToken)
	if err != nil {
		return
	}

	output = GetAuthByRefreshTokenOutputDTO{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		Active:                auth.GetActive(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	}

	return
}
