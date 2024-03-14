package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(body io.ReadCloser) (err error)
	GetUser(id int32) (output usecase.GetUserOutputDTO, err error)
	GetUsers(limit int32, offset int32) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(id int32, body io.ReadCloser) (err error)
	DeleteUser(id int32) (err error)
	Register(body io.ReadCloser) (output usecase.RegisterOutputDTO, err error)
	UserVerify(body io.ReadCloser) (err error)
	ForgotPassword(body io.ReadCloser) (err error)
	UpdatePasswordUser(body io.ReadCloser) (err error)
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(body io.ReadCloser) (err error) {
	s.userRepo.Begin()

	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, _ := usecase.NewGetUserByEmailUseCase().Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if output.ID != 0 {
		slog.Info("email already exists")
		return errors.New("email already exists")
	}

	if err = usecase.NewCreateUserUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("User created")
	return
}

func (s *UserService) GetUser(id int32) (output usecase.GetUserOutputDTO, err error) {

	input := usecase.GetUserInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetUserUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User found")
	return
}

func (s *UserService) GetUsers(limit int32, offset int32) (output []usecase.GetUsersOutputDTO, err error) {

	input := usecase.GetUsersInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	output, err = usecase.NewGetUsersUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Users found")
	return
}

func (s *UserService) UpdateUser(id int32, body io.ReadCloser) (err error) {
	s.userRepo.Begin()
	input := usecase.UpdateUserInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	if err = usecase.NewUpdateUserUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("User updated")
	return
}

func (s *UserService) DeleteUser(id int32) (err error) {
	s.userRepo.Begin()
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteUserUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("User deleted")
	return
}

func (s *UserService) Register(body io.ReadCloser) (output usecase.RegisterOutputDTO, err error) {

	s.userRepo.Begin()

	input := usecase.RegisterInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := usecase.NewGetUserByEmailUseCase().Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if check.ID != 0 {
		slog.Info("email already exists")
		return usecase.RegisterOutputDTO{}, errors.New("email already exists")
	}

	output, err = usecase.NewRegisterUseCase().Execute(input)
	if err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}

	s.userRepo.Commit()
	slog.Info("User registered")
	return
}

func (s *UserService) UserVerify(body io.ReadCloser) (err error) {
	s.userRepo.Begin()
	input := usecase.UserVerifyInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUserVerifyUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("User verified")
	return
}

func (s *UserService) ForgotPassword(body io.ReadCloser) (err error) {
	s.userRepo.Begin()
	input := usecase.ForgotPasswordInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewForgotPasswordUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("Email Sended")
	return
}

func (s *UserService) UpdatePasswordUser(body io.ReadCloser) (err error) {
	s.userRepo.Begin()
	input := usecase.UpdatePasswordUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdatePasswordUserUseCase().Execute(input); err != nil {
		s.userRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.userRepo.Commit()
	slog.Info("Password updated")
	return
}
