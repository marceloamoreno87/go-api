package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetAuthByUserIDInputDTO struct {
	UserID int32 `json:"user_id"`
}

type GetAuthByUserIDOutputDTO struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}

type GetAuthByUserIDUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewGetAuthByUserIDUseCase(DB config.SQLCInterface) *GetAuthByUserIDUseCase {
	return &GetAuthByUserIDUseCase{
		repo: repository.NewAuthRepository(DB),
	}
}

func (uc *GetAuthByUserIDUseCase) Execute(input GetAuthByUserIDInputDTO) (output GetAuthByUserIDOutputDTO, err error) {
	auth, err := uc.repo.GetAuthByUserID(input.UserID)
	if err != nil {
		return
	}

	output = GetAuthByUserIDOutputDTO{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		Active:                auth.GetActive(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	}

	return
}
