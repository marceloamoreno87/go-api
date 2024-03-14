package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(body io.ReadCloser) (err error)
	GetUserById(id int32) (output usecase.GetUserOutputDTO, err error)
	GetUserByEmail(email string) (output usecase.GetUserOutputDTO, err error)
	GetUsers(limit int32, offset int32) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(id int32, body io.ReadCloser) (err error)
	DeleteUser(id int32) (err error)
	UpdateUserPassword(id int32, body io.ReadCloser) (err error)
	config.SQLCInterface
}

type UserService struct {
	config.SQLCInterface
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(body io.ReadCloser) (err error) {
	s.Begin()

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
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("User created")
	return
}

func (s *UserService) GetUserById(id int32) (output usecase.GetUserOutputDTO, err error) {

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

func (s *UserService) GetUserByEmail(email string) (output usecase.GetUserOutputDTO, err error) {
	input := usecase.GetUserByEmailInputDTO{
		Email: email,
	}
	output, err = usecase.NewGetUserByEmailUseCase().Execute(input)
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
	s.Begin()
	input := usecase.UpdateUserInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	if err = usecase.NewUpdateUserUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("User updated")
	return
}

func (s *UserService) DeleteUser(id int32) (err error) {
	s.Begin()
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteUserUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("User deleted")
	return
}

func (s *UserService) UpdateUserPassword(id int32, body io.ReadCloser) (err error) {
}
