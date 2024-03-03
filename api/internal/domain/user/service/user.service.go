package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(body io.ReadCloser) (err error)
	GetUser(id int32) (output usecase.GetUserOutputDTO, err error)
	GetUsers(limit int32, offset int32) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(id int32, body io.ReadCloser) (err error)
	DeleteUser(id int32) (err error)
	Login(body io.ReadCloser) (output usecase.LoginOutputDTO, err error)
	Register(body io.ReadCloser) (output usecase.RegisterOutputDTO, err error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(body io.ReadCloser) (err error) {
	s.repo.Begin()

	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, _ := usecase.NewGetUserByEmailUseCase(s.repo).Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if output.ID != 0 {
		slog.Info("email already exists")
		return errors.New("email already exists")
	}

	if err = usecase.NewCreateUserUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("User created")
	return
}

func (s *UserService) GetUser(id int32) (output usecase.GetUserOutputDTO, err error) {

	input := usecase.GetUserInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetUserUseCase(s.repo).Execute(input)
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

	output, err = usecase.NewGetUsersUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Users found")
	return
}

func (s *UserService) UpdateUser(id int32, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdateUserInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	if err = usecase.NewUpdateUserUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("User updated")
	return
}

func (s *UserService) DeleteUser(id int32) (err error) {
	s.repo.Begin()
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteUserUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("User deleted")
	return
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
