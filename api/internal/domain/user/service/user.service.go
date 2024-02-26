package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(body io.ReadCloser) (err error) {

	input := usecase.CreateUserInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewCreateUserUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *UserService) GetUser(id string) (output usecase.GetUserOutputDTO, err error) {

	input := usecase.GetUserInputDTO{
		ID: helper.StrToInt32(id),
	}

	output, err = usecase.NewGetUserUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *UserService) GetUsers(limit string, offset string) (output []usecase.GetUsersOutputDTO, err error) {

	input := usecase.GetUsersInputDTO{
		Limit:  helper.StrToInt32(limit),
		Offset: helper.StrToInt32(offset),
	}

	output, err = usecase.NewGetUsersUseCase(s.repo).Execute(input)

	return
}

func (s *UserService) UpdateUser(id string, body io.ReadCloser) (err error) {

	input := usecase.UpdateUserInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewUpdateUserUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *UserService) DeleteUser(id string) (err error) {

	input := usecase.DeleteUserInputDTO{
		ID: helper.StrToInt32(id),
	}

	err = usecase.NewDeleteUserUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}
