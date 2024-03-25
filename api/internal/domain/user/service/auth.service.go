package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/response"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthService struct {
	db                             config.SQLCInterface
	GetUserUseCase                 usecase.GetUserUseCase
	GetUserByEmailUseCase          usecase.GetUserByEmailUseCase
	CreateAuthUseCase              usecase.CreateAuthUseCase
	LoginUserUseCase               usecase.LoginUserUseCase
	GetAuthByRefreshTokenUseCase   usecase.GetAuthByRefreshTokenUseCase
	UpdateAuthRevokeUseCase        usecase.UpdateAuthRevokeUseCase
	GetAuthByTokenUseCase          usecase.GetAuthByTokenUseCase
	GetAuthByUserIDUseCase         usecase.GetAuthByUserIDUseCase
	CreateUserUseCase              usecase.CreateUserUseCase
	UpdateUserPasswordUseCase      usecase.UpdateUserPasswordUseCase
	GetUserValidationByHashUseCase usecase.GetUserValidationByHashUseCase
	UpdateUserActive               usecase.UpdateUserActiveUseCase
	UpdateUserValidationUsed       usecase.UpdateUserValidationUsedUseCase
	CreateUserValidationUseCase    usecase.CreateUserValidationUseCase
}

func NewAuthService() *AuthService {
	db := config.NewSqlc(config.DB)
	return &AuthService{
		db:                             db,
		GetUserUseCase:                 *usecase.NewGetUserUseCase(db),
		GetUserByEmailUseCase:          *usecase.NewGetUserByEmailUseCase(db),
		CreateAuthUseCase:              *usecase.NewCreateAuthUseCase(db),
		LoginUserUseCase:               *usecase.NewLoginUserUseCase(),
		GetAuthByRefreshTokenUseCase:   *usecase.NewGetAuthByRefreshTokenUseCase(db),
		UpdateAuthRevokeUseCase:        *usecase.NewUpdateAuthRevokeUseCase(db),
		GetAuthByTokenUseCase:          *usecase.NewGetAuthByTokenUseCase(db),
		GetAuthByUserIDUseCase:         *usecase.NewGetAuthByUserIDUseCase(db),
		CreateUserUseCase:              *usecase.NewCreateUserUseCase(db),
		UpdateUserPasswordUseCase:      *usecase.NewUpdateUserPasswordUseCase(db),
		GetUserValidationByHashUseCase: *usecase.NewGetUserValidationByHashUseCase(db),
		UpdateUserActive:               *usecase.NewUpdateUserActiveUseCase(db),
		UpdateUserValidationUsed:       *usecase.NewUpdateUserValidationUsedUseCase(db),
		CreateUserValidationUseCase:    *usecase.NewCreateUserValidationUseCase(db),
	}
}

func (s *AuthService) Login(ctx context.Context, input request.RequestLogin) (output response.ResponseLogin, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	user, err := s.GetUserByEmailUseCase.Execute(ctx, usecase.GetUserByEmailInputDTO{Email: input.Email})
	if err != nil {
		slog.Info("err", err)
		return
	}

	_, err = s.LoginUserUseCase.Execute(usecase.LoginUserInputDTO{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		RoleID:          user.RoleID,
		AvatarID:        user.AvatarID,
		RequestPassword: input.Password,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	auth, _ := s.GetAuthByUserIDUseCase.Execute(ctx, usecase.GetAuthByUserIDInputDTO{
		UserID: user.ID,
	})
	if auth.UserID != 0 {
		return response.ResponseLogin{
			Token:                 auth.Token,
			RefreshToken:          auth.RefreshToken,
			UserID:                auth.UserID,
			Active:                auth.Active,
			TokenExpiresIn:        auth.TokenExpiresIn,
			RefreshTokenExpiresIn: auth.RefreshTokenExpiresIn,
		}, nil
	}

	_, err = s.UpdateAuthRevokeUseCase.Execute(ctx, usecase.UpdateAuthRevokeInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	newToken, err := s.CreateAuthUseCase.Execute(ctx, usecase.CreateAuthInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	output = response.ResponseLogin{
		Token:                 newToken.Token,
		RefreshToken:          newToken.RefreshToken,
		UserID:                newToken.UserID,
		Active:                newToken.Active,
		TokenExpiresIn:        newToken.TokenExpiresIn,
		RefreshTokenExpiresIn: newToken.RefreshTokenExpiresIn,
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("User logged in")
	return
}

func (s *AuthService) RefreshToken(ctx context.Context, input request.RequestRefreshToken) (output response.ResponseRefreshToken, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	rt, err := s.GetAuthByRefreshTokenUseCase.Execute(ctx, usecase.GetAuthByRefreshTokenInputDTO{
		UserID:       input.UserID,
		RefreshToken: input.RefreshToken,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	_, err = s.UpdateAuthRevokeUseCase.Execute(ctx, usecase.UpdateAuthRevokeInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	token, err := s.CreateAuthUseCase.Execute(ctx, usecase.CreateAuthInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	output = response.ResponseRefreshToken{
		Token:                 token.Token,
		RefreshToken:          token.RefreshToken,
		UserID:                token.UserID,
		Active:                token.Active,
		TokenExpiresIn:        token.TokenExpiresIn,
		RefreshTokenExpiresIn: token.RefreshTokenExpiresIn,
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("Token refreshed")
	return
}
