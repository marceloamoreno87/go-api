package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserService struct {
	DB config.SQLCInterface
}

func NewUserService() *UserService {
	return &UserService{
		DB: config.Sqcl,
	}
}

func (s *UserService) CreateUser(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error) {
	s.DB.Begin()

	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := usecase.NewGetUserByEmailUseCase().Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if check.ID != 0 {
		slog.Info("email already exists")
		return output, errors.New("email already exists")
	}

	output, err = usecase.NewCreateUserUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
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

func (s *UserService) GetUserByEmail(email string) (output usecase.GetUserByEmailOutputDTO, err error) {
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

func (s *UserService) UpdateUser(id int32, body io.ReadCloser) (output usecase.UpdateUserOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.UpdateUserInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewUpdateUserUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("User updated")
	return
}

func (s *UserService) DeleteUser(id int32) (output usecase.DeleteUserOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	output, err = usecase.NewDeleteUserUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("User deleted")
	return
}

func (s *UserService) UpdateUserPassword(id int32, body io.ReadCloser) (output usecase.UpdateUserPasswordOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.UpdateUserPasswordInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewUpdateUserPasswordUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("User password updated")
	return
}
