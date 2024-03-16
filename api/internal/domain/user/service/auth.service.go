package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"time"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthService struct {
	NewGetUserUseCase                 usecaseInterface.GetUserUseCaseInterface
	NewGetUserByEmailUseCase          usecaseInterface.GetUserByEmailUseCaseInterface
	NewCreateAuthUseCase              usecaseInterface.CreateAuthUseCaseInterface
	NewLoginUserUseCase               usecaseInterface.LoginUserUseCaseInterface
	NewGetAuthByRefreshTokenUseCase   usecaseInterface.GetAuthByRefreshTokenUseCase
	NewUpdateAuthRevokeUseCase        usecaseInterface.UpdateAuthRevokeUseCaseInterface
	NewGetAuthByTokenUseCase          usecaseInterface.GetAuthByTokenUseCaseInterface
	NewGetAuthByUserIDUseCase         usecaseInterface.GetAuthByUserIDUseCaseInterface
	NewCheckTokenUseCase              usecaseInterface.CheckTokenUseCaseInterface
	NewCheckRefreshTokenUseCase       usecaseInterface.CheckRefreshTokenUseCaseInterface
	NewCreateUserUseCase              usecaseInterface.CreateUserUseCaseInterface
	NewUpdateUserPasswordUseCase      usecaseInterface.UpdateUserPasswordUseCaseInterface
	NewGetUserValidationByHashUseCase usecaseInterface.GetUserValidationByHashUseCaseInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		NewGetUserUseCase:                 usecase.NewGetUserUseCase(),
		NewGetUserByEmailUseCase:          usecase.NewGetUserByEmailUseCase(),
		NewCreateAuthUseCase:              usecase.NewCreateAuthUseCase(),
		NewLoginUserUseCase:               usecase.NewLoginUserUseCase(),
		NewGetAuthByRefreshTokenUseCase:   usecase.NewGetAuthByRefreshTokenUseCase(),
		NewUpdateAuthRevokeUseCase:        usecase.NewUpdateAuthRevokeUseCase(),
		NewGetAuthByTokenUseCase:          usecase.NewGetAuthByTokenUseCase(),
		NewGetAuthByUserIDUseCase:         usecase.NewGetAuthByUserIDUseCase(),
		NewCheckTokenUseCase:              usecase.NewCheckTokenUseCase(),
		NewCheckRefreshTokenUseCase:       usecase.NewCheckRefreshTokenUseCase(),
		NewCreateUserUseCase:              usecase.NewCreateUserUseCase(),
		NewUpdateUserPasswordUseCase:      usecase.NewUpdateUserPasswordUseCase(),
		NewGetUserValidationByHashUseCase: usecase.NewGetUserValidationByHashUseCase(),
	}
}

func (s *AuthService) Login(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error) {

	input := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	// get user by email
	user, err := s.NewGetUserByEmailUseCase.Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if err != nil {
		slog.Info("err", err)
		return
	}

	logged, err := s.NewLoginUserUseCase.Execute(usecase.LoginUserInputDTO{
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

	if !logged.Valid {
		slog.Info("Invalid user")
		return usecase.CreateAuthOutputDTO{}, errors.New("invalid user")
	}

	// consulta se tem token valido
	token, _ := s.NewGetAuthByUserIDUseCase.Execute(usecase.GetAuthByUserIDInputDTO{
		UserID: user.ID,
	})

	check, err := s.NewCheckTokenUseCase.Execute(usecase.CheckTokenInputDTO{
		UserID: user.ID,
		Token:  token.Token,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	if check.Valid {
		output = usecase.CreateAuthOutputDTO{
			Token:                 token.Token,
			RefreshToken:          token.RefreshToken,
			UserID:                token.UserID,
			Active:                token.Active,
			TokenExpiresIn:        token.TokenExpiresIn,
			RefreshTokenExpiresIn: token.RefreshTokenExpiresIn,
		}
		return
	}

	_, err = s.NewUpdateAuthRevokeUseCase.Execute(usecase.UpdateAuthRevokeInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	newToken, err := s.NewCreateAuthUseCase.Execute(usecase.CreateAuthInputDTO{
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

func (s *AuthService) RefreshToken(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error) {

	input := struct {
		UserID       int32  `json:"user_id"`
		RefreshToken string `json:"refresh_token"`
	}{}

	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	// Consultar o refresh token e user ID no banco
	rt, err := s.NewGetAuthByRefreshTokenUseCase.Execute(usecase.GetAuthByRefreshTokenInputDTO{
		UserID:       input.UserID,
		RefreshToken: input.RefreshToken,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	// Validar se o token é válido - Se for válido não gera novo token
	checkToken, err := s.NewCheckTokenUseCase.Execute(usecase.CheckTokenInputDTO{
		UserID: rt.UserID,
		Token:  rt.Token,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	if checkToken.Valid {
		return output, errors.New("token is valid")
	}

	checkRefreshToken, err := s.NewCheckRefreshTokenUseCase.Execute(usecase.CheckRefreshTokenInputDTO{
		UserID:       rt.UserID,
		RefreshToken: rt.RefreshToken,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	if !checkRefreshToken.Valid {
		return output, errors.New("invalid refresh token")
	}

	_, err = s.NewUpdateAuthRevokeUseCase.Execute(usecase.UpdateAuthRevokeInputDTO{
		UserID: rt.UserID,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	token, err := s.NewCreateAuthUseCase.Execute(usecase.CreateAuthInputDTO{
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

func (s *AuthService) Register(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error) {
	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = s.NewCreateUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User registered")
	return
}

func (s *AuthService) UpdateUserPassword(body io.ReadCloser) (output usecase.UpdateUserPasswordOutputDTO, err error) {

	input := struct {
		Hash     string `json:"hash"`
		Password string `json:"password"`
	}{}

	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	userValidation, err := s.NewGetUserValidationByHashUseCase.Execute(usecase.GetUserValidationByHashInputDTO{
		Hash: input.Hash,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	if userValidation.Used {
		return output, errors.New("hash already used")
	}

	if int64(userValidation.ExpiresIn) > time.Now().Unix() {
		return output, errors.New("hash expired")
	}

	user, err := s.NewGetUserUseCase.Execute(usecase.GetUserInputDTO{ID: userValidation.UserID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	//atualizar senha
	output, err = s.NewUpdateUserPasswordUseCase.Execute(usecase.UpdateUserPasswordInputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: input.Password,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User password updated")
	return
}
