package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
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
	repo repository.Authrepository
}

func NewCreateAuthUseCase(db config.SQLCInterface) *CreateAuthUseCase {
	return &CreateAuthUseCase{
		repo: repository.NewAuthRepository(db),
	}
}

func (uc *CreateAuthUseCase) Execute(ctx context.Context, input CreateAuthInputDTO) (output CreateAuthOutputDTO, err error) {
	auth, err := entity.NewAuth(input.UserID)
	if err != nil {
		return
	}

	err = uc.repo.CreateAuth(ctx, auth)
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
