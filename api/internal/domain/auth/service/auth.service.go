package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/repository"
	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
)

type AuthServiceInterface interface {
	Login(body io.ReadCloser) (output usecase.LoginOutputDTO, err error)
	RefreshToken(body io.ReadCloser) (output usecase.RefreshTokenOutputDTO, err error)
}

type AuthService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(
	repo repository.AuthRepositoryInterface,
) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(body io.ReadCloser) (output usecase.LoginOutputDTO, err error) {
	input := usecase.LoginInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewLoginUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User logged in")
	return
}

func (s *AuthService) RefreshToken(body io.ReadCloser) (output usecase.RefreshTokenOutputDTO, err error) {
	input := usecase.RefreshTokenInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewRefreshTokenUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Token refreshed")
	return
}
