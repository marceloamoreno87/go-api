package service

import (
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthService struct {
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
	return &AuthService{
		GetUserUseCase:                 usecase.NewGetUserUseCase(),
		GetUserByEmailUseCase:          usecase.NewGetUserByEmailUseCase(),
		CreateAuthUseCase:              usecase.NewCreateAuthUseCase(),
		LoginUserUseCase:               usecase.NewLoginUserUseCase(),
		GetAuthByRefreshTokenUseCase:   usecase.NewGetAuthByRefreshTokenUseCase(),
		UpdateAuthRevokeUseCase:        usecase.NewUpdateAuthRevokeUseCase(),
		GetAuthByTokenUseCase:          usecase.NewGetAuthByTokenUseCase(),
		GetAuthByUserIDUseCase:         usecase.NewGetAuthByUserIDUseCase(),
		CreateUserUseCase:              usecase.NewCreateUserUseCase(),
		UpdateUserPasswordUseCase:      usecase.NewUpdateUserPasswordUseCase(),
		GetUserValidationByHashUseCase: usecase.NewGetUserValidationByHashUseCase(),
		UpdateUserActive:               usecase.NewUpdateUserActiveUseCase(),
		UpdateUserValidationUsed:       usecase.NewUpdateUserValidationUsedUseCase(),
		CreateUserValidationUseCase:    usecase.NewCreateUserValidationUseCase(),
	}
}

func (s *AuthService) Login(input request.RequestLoginInputDTO) (output usecase.CreateAuthOutputDTO, err error) {
	user, err := s.GetUserByEmailUseCase.Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
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

	auth, _ := s.GetAuthByUserIDUseCase.Execute(usecase.GetAuthByUserIDInputDTO{
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

	_, err = s.UpdateAuthRevokeUseCase.Execute(usecase.UpdateAuthRevokeInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	newToken, err := s.CreateAuthUseCase.Execute(usecase.CreateAuthInputDTO{
		UserID: user.ID,
	})
	if err != nil {
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
	slog.Info("User logged in")
	return
}

func (s *AuthService) RefreshToken(input request.RequestRefreshTokenInputDTO) (output usecase.CreateAuthOutputDTO, err error) {
	rt, err := s.GetAuthByRefreshTokenUseCase.Execute(usecase.GetAuthByRefreshTokenInputDTO{
		UserID:       input.UserID,
		RefreshToken: input.RefreshToken,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	_, err = s.UpdateAuthRevokeUseCase.Execute(usecase.UpdateAuthRevokeInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	token, err := s.CreateAuthUseCase.Execute(usecase.CreateAuthInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
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
	slog.Info("Token refreshed")
	return
}
