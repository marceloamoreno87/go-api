package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error)
	// RefreshToken(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error)
	config.SQLCInterface
}

type AuthService struct {
	config.SQLCInterface
}

func NewAuthService() *AuthService {
	return &AuthService{}
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

// func (s *AuthService) RefreshToken(body io.ReadCloser) (output usecase.RefreshTokenOutputDTO, err error) {
// 	input := usecase.RefreshTokenInputDTO{}
// 	if err = json.NewDecoder(body).Decode(&input); err != nil {
// 		slog.Info("err", err)
// 		return
// 	}
// 	output, err = usecase.NewRefreshTokenUseCase().Execute(input)
// 	if err != nil {
// 		slog.Info("err", err)
// 		return
// 	}
// 	slog.Info("Token refreshed")
// 	return
// }