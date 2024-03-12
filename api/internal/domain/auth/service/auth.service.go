package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/repository"
	"github.com/marceloamoreno/goapi/internal/domain/auth/usecase"
)

type UserServiceInterface interface {
	Login(body io.ReadCloser) (output usecase.LoginOutputDTO, err error)
	Register(body io.ReadCloser) (output usecase.RegisterOutputDTO, err error)
	UserVerify(body io.ReadCloser) (err error)
	ForgotPassword(body io.ReadCloser) (err error)
	UpdatePasswordUser(body io.ReadCloser) (err error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(
	repo repository.UserRepositoryInterface,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(body io.ReadCloser) (output usecase.LoginOutputDTO, err error) {
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

func (s *UserService) Register(body io.ReadCloser) (output usecase.RegisterOutputDTO, err error) {

	s.repo.Begin()

	input := usecase.RegisterInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := usecase.NewGetUserByEmailUseCase(s.repo).Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if check.ID != 0 {
		slog.Info("email already exists")
		return usecase.RegisterOutputDTO{}, errors.New("email already exists")
	}

	output, err = usecase.NewRegisterUseCase(s.repo).Execute(input)
	if err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}

	s.repo.Commit()
	slog.Info("User registered")
	return
}

func (s *UserService) UserVerify(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UserVerifyInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUserVerifyUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("User verified")
	return
}

func (s *UserService) ForgotPassword(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.ForgotPasswordInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewForgotPasswordUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Email Sended")
	return
}

func (s *UserService) UpdatePasswordUser(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdatePasswordUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdatePasswordUserUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Password updated")
	return
}
