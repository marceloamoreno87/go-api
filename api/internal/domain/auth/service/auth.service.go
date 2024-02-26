package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type AuthService struct {
	repo repository.UserRepositoryInterface
}

func NewAuthService(repo repository.UserRepositoryInterface) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(body io.ReadCloser) (output usecase.GetJWTOutputDTO, err error) {

	input := usecase.GetJWTInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewGetJWTUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *AuthService) Refresh(body io.ReadCloser) (output usecase.GetRefreshJWTOutputDTO, err error) {

	input := usecase.GetRefreshJWTInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewGetRefreshJWTUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}
