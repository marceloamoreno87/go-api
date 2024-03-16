package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserService struct {
	NewGetUserByEmailUseCase     usecaseInterface.GetUserByEmailUseCaseInterface
	NewCreateUserUseCase         usecaseInterface.CreateUserUseCaseInterface
	NewGetUserUseCase            usecaseInterface.GetUserUseCaseInterface
	NewGetUsersUseCase           usecaseInterface.GetUsersUseCaseInterface
	NewUpdateUserUseCase         usecaseInterface.UpdateUserUseCaseInterface
	NewDeleteUserUseCase         usecaseInterface.DeleteUserUseCaseInterface
	NewUpdateUserPasswordUseCase usecaseInterface.UpdateUserPasswordUseCaseInterface
}

func NewUserService() *UserService {
	return &UserService{
		NewGetUserByEmailUseCase:     usecase.NewGetUserByEmailUseCase(),
		NewCreateUserUseCase:         usecase.NewCreateUserUseCase(),
		NewGetUserUseCase:            usecase.NewGetUserUseCase(),
		NewGetUsersUseCase:           usecase.NewGetUsersUseCase(),
		NewUpdateUserUseCase:         usecase.NewUpdateUserUseCase(),
		NewDeleteUserUseCase:         usecase.NewDeleteUserUseCase(),
		NewUpdateUserPasswordUseCase: usecase.NewUpdateUserPasswordUseCase(),
	}
}

func (s *UserService) CreateUser(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error) {
	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := s.NewGetUserByEmailUseCase.Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if check.ID != 0 {
		slog.Info("email already exists")
		return output, errors.New("email already exists")
	}

	output, err = s.NewCreateUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User created")
	return
}

func (s *UserService) GetUserById(id int32) (output usecase.GetUserOutputDTO, err error) {
	input := usecase.GetUserInputDTO{
		ID: id,
	}

	output, err = s.NewGetUserUseCase.Execute(input)
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
	output, err = s.NewGetUserByEmailUseCase.Execute(input)
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

	output, err = s.NewGetUsersUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Users found")
	return
}

func (s *UserService) UpdateUser(id int32, body io.ReadCloser) (output usecase.UpdateUserOutputDTO, err error) {
	input := usecase.UpdateUserInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = s.NewUpdateUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User updated")
	return
}

func (s *UserService) DeleteUser(id int32) (output usecase.DeleteUserOutputDTO, err error) {
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	output, err = s.NewDeleteUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User deleted")
	return
}

func (s *UserService) UpdateUserPassword(id int32, body io.ReadCloser) (output usecase.UpdateUserPasswordOutputDTO, err error) {
	input := usecase.UpdateUserPasswordInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = s.NewUpdateUserPasswordUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User password updated")
	return
}
