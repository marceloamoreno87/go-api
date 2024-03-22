package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AuthRepository struct {
	db config.SQLCInterface
}

func NewAuthRepository(db config.SQLCInterface) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) CreateAuth(ctx context.Context, auth entityInterface.AuthInterface) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateAuth(ctx, db.CreateAuthParams{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	})
}

func (repo *AuthRepository) GetAuthByUserID(ctx context.Context, userId int32) (output entityInterface.AuthInterface, err error) {
	a, err := repo.db.GetDbQueries().GetAuthByUserID(ctx, userId)
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) UpdateAuthRevokeByUserID(ctx context.Context, userId int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateAuthRevokeByUserID(ctx, userId)
}

func (repo *AuthRepository) GetAuthByToken(ctx context.Context, userId int32, token string) (output entityInterface.AuthInterface, err error) {
	a, err := repo.db.GetDbQueries().GetAuthByToken(ctx, db.GetAuthByTokenParams{
		Token:  token,
		UserID: userId,
	})
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) GetAuthByRefreshToken(ctx context.Context, userId int32, refreshToken string) (output entityInterface.AuthInterface, err error) {
	a, err := repo.db.GetDbQueries().GetAuthByRefreshToken(ctx, db.GetAuthByRefreshTokenParams{
		RefreshToken: refreshToken,
		UserID:       userId,
	})
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}
