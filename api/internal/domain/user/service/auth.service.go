package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthService struct {
	db                             config.SQLCInterface
	GetUserUseCase                 usecaseInterface.GetUserUseCaseInterface
	GetUserByEmailUseCase          usecaseInterface.GetUserByEmailUseCaseInterface
	CreateAuthUseCase              usecaseInterface.CreateAuthUseCaseInterface
	LoginUserUseCase               usecaseInterface.LoginUserUseCaseInterface
	GetAuthByRefreshTokenUseCase   usecaseInterface.GetAuthByRefreshTokenUseCase
	UpdateAuthRevokeUseCase        usecaseInterface.UpdateAuthRevokeUseCaseInterface
	GetAuthByTokenUseCase          usecaseInterface.GetAuthByTokenUseCaseInterface
	GetAuthByUserIDUseCase         usecaseInterface.GetAuthByUserIDUseCaseInterface
	CreateUserUseCase              usecaseInterface.CreateUserUseCaseInterface
	UpdateUserPasswordUseCase      usecaseInterface.UpdateUserPasswordUseCaseInterface
	GetUserValidationByHashUseCase usecaseInterface.GetUserValidationByHashUseCaseInterface
	UpdateUserActive               usecaseInterface.UpdateUserActiveUseCaseInterface
	UpdateUserValidationUsed       usecaseInterface.UpdateUserValidationUsedUseCaseInterface
	CreateUserValidationUseCase    usecaseInterface.CreateUserValidationUseCaseInterface
}

func NewAuthService() *AuthService {
	db := config.NewSqlc(config.DB)
	return &AuthService{
		db:                             db,
		GetUserUseCase:                 usecase.NewGetUserUseCase(db),
		GetUserByEmailUseCase:          usecase.NewGetUserByEmailUseCase(db),
		CreateAuthUseCase:              usecase.NewCreateAuthUseCase(db),
		LoginUserUseCase:               usecase.NewLoginUserUseCase(),
		GetAuthByRefreshTokenUseCase:   usecase.NewGetAuthByRefreshTokenUseCase(db),
		UpdateAuthRevokeUseCase:        usecase.NewUpdateAuthRevokeUseCase(db),
		GetAuthByTokenUseCase:          usecase.NewGetAuthByTokenUseCase(db),
		GetAuthByUserIDUseCase:         usecase.NewGetAuthByUserIDUseCase(db),
		CreateUserUseCase:              usecase.NewCreateUserUseCase(db),
		UpdateUserPasswordUseCase:      usecase.NewUpdateUserPasswordUseCase(db),
		GetUserValidationByHashUseCase: usecase.NewGetUserValidationByHashUseCase(db),
		UpdateUserActive:               usecase.NewUpdateUserActiveUseCase(db),
		UpdateUserValidationUsed:       usecase.NewUpdateUserValidationUsedUseCase(db),
		CreateUserValidationUseCase:    usecase.NewCreateUserValidationUseCase(db),
	}
}

func (s *AuthService) Login(ctx context.Context, input request.RequestLogin) (output usecase.CreateAuthOutputDTO, err error) {
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
		return usecase.CreateAuthOutputDTO{
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
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	newToken, err := s.CreateAuthUseCase.Execute(ctx, usecase.CreateAuthInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	output = usecase.CreateAuthOutputDTO{
		Token:                 newToken.Token,
		RefreshToken:          newToken.RefreshToken,
		UserID:                newToken.UserID,
		Active:                newToken.Active,
		TokenExpiresIn:        newToken.TokenExpiresIn,
		RefreshTokenExpiresIn: newToken.RefreshTokenExpiresIn,
	}
	tx.Commit()
	slog.Info("User logged in")
	return
}

func (s *AuthService) RefreshToken(ctx context.Context, input request.RequestRefreshToken) (output usecase.CreateAuthOutputDTO, err error) {
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
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	token, err := s.CreateAuthUseCase.Execute(ctx, usecase.CreateAuthInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	output = usecase.CreateAuthOutputDTO{
		Token:                 token.Token,
		RefreshToken:          token.RefreshToken,
		UserID:                token.UserID,
		Active:                token.Active,
		TokenExpiresIn:        token.TokenExpiresIn,
		RefreshTokenExpiresIn: token.RefreshTokenExpiresIn,
	}
	tx.Commit()
	slog.Info("Token refreshed")
	return
}
