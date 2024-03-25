package usecase

import (
	"context"
	"errors"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
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
	repo repository.Authrepository
}

func NewGetAuthByRefreshTokenUseCase(db config.SQLCInterface) *GetAuthByRefreshTokenUseCase {
	return &GetAuthByRefreshTokenUseCase{
		repo: repository.NewAuthRepository(db),
	}
}

func (uc *GetAuthByRefreshTokenUseCase) Execute(ctx context.Context, input GetAuthByRefreshTokenInputDTO) (output GetAuthByRefreshTokenOutputDTO, err error) {
	auth, err := uc.repo.GetAuthByRefreshToken(ctx, input.UserID, input.RefreshToken)
	if err != nil {
		return
	}

	a, err := entity.NewAuth(input.UserID)
	if err != nil {
		return
	}
	a.SetToken(auth.GetToken())

	if valid, _ := a.IsValidToken(); valid {
		return output, errors.New("token is valid")
	}

	if valid, _ := a.IsValidRefreshToken(); !valid {
		return output, errors.New("invalid refresh token")
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
