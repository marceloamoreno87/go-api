package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetAuthByTokenInputDTO struct {
	ID    int32  `json:"id"`
	Token string `json:"token"`
}

type GetAuthByTokenOutputDTO struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}

type GetAuthByTokenUseCase struct {
	repo repository.Authrepository
}

func NewGetAuthByTokenUseCase(db config.SQLCInterface) *GetAuthByTokenUseCase {
	return &GetAuthByTokenUseCase{
		repo: repository.NewAuthRepository(db),
	}
}

func (uc *GetAuthByTokenUseCase) Execute(ctx context.Context, input GetAuthByTokenInputDTO) (output GetAuthByTokenOutputDTO, err error) {
	auth, err := uc.repo.GetAuthByToken(ctx, input.ID, input.Token)
	if err != nil {
		return
	}

	output = GetAuthByTokenOutputDTO{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		Active:                auth.GetActive(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	}

	return
}
