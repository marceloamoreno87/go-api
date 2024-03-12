package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
	"github.com/marceloamoreno/goapi/internal/shared/repository"
)

type AuthRepositoryInterface interface {
	CreateToken(auth *entity.Auth) (err error)
	GetTokenByUser() (auth *entity.Auth, err error)
	RevokeTokenByUser(*entity.Auth) error
	repository.RepositoryInterface
}

type AuthRepository struct {
	repository.Repository
}

func NewAuthRepository(DB config.DatabaseInterface) *AuthRepository {
	return &AuthRepository{
		Repository: *repository.NewRepository(DB),
	}
}

func (repo *AuthRepository) CreateToken(auth *entity.Auth) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateToken(context.Background(), db.CreateTokenParams{
		UserID:       auth.UserID,
		Token:        auth.Token,
		RefreshToken: auth.RefreshToken,
		ExpiresIn:    auth.ExpiresIn,
	})
	return
}

func (repo *AuthRepository) GetTokenByUser() (auth *entity.Auth, err error) {
	auth, err = repo.GetDbQueries().GetTokenByUser(context.Background(), db.GetTokenByUserParams{
		UserID: auth.UserID,
	})
	return
}

func (repo *AuthRepository) RevokeTokenByUser(auth *entity.Auth) error {
	err := repo.GetDbQueries().WithTx(repo.GetTx()).RevokeTokenByUser(context.Background(), db.RevokeTokenByUserParams{
		UserID: auth.UserID,
	})
	return err
}
