package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthService struct {
	DB config.SQLCInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		DB: config.Sqcl,
	}
}

func (s *AuthService) Login(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error) {
	input := usecase.CreateAuthInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewAuthUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User logged in")
	return
}

// TODO: REFACTOR
func (s *AuthService) RefreshToken(body io.ReadCloser) (output usecase.CreateAuthByRefreshTokenOutputDTO, err error) {
	input := usecase.CreateAuthByRefreshTokenInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewCreateAuthByRefreshTokenUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User logged in")
	return
}

func (s *AuthService) Register(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error) {
	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewCreateUserUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User registered")
	return
}

func (s *AuthService) UpdateUserPassword(body io.ReadCloser) (output usecase.UpdateUserPasswordOutputDTO, err error) {
	input := usecase.UpdateUserPasswordInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewUpdateUserPasswordUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User password updated")
	return
}
