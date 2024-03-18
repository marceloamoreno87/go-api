package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RequestUpdateUserPasswordInputDTO struct {
	Hash     string `json:"hash"`
	Password string `json:"password"`
}

type UserService struct {
	GetUserByEmailUseCase          usecaseInterface.GetUserByEmailUseCaseInterface
	CreateUserUseCase              usecaseInterface.CreateUserUseCaseInterface
	GetUserUseCase                 usecaseInterface.GetUserUseCaseInterface
	GetUsersUseCase                usecaseInterface.GetUsersUseCaseInterface
	UpdateUserUseCase              usecaseInterface.UpdateUserUseCaseInterface
	DeleteUserUseCase              usecaseInterface.DeleteUserUseCaseInterface
	UpdateUserPasswordUseCase      usecaseInterface.UpdateUserPasswordUseCaseInterface
	CreateUserValidationUseCase    usecaseInterface.CreateUserValidationUseCaseInterface
	GetUserValidationByHashUseCase usecaseInterface.GetUserValidationByHashUseCaseInterface
	UpdateUserValidationUsed       usecaseInterface.UpdateUserValidationUsedUseCaseInterface
}

func NewUserService() *UserService {
	return &UserService{
		GetUserByEmailUseCase:          usecase.NewGetUserByEmailUseCase(),
		CreateUserUseCase:              usecase.NewCreateUserUseCase(),
		GetUserUseCase:                 usecase.NewGetUserUseCase(),
		GetUsersUseCase:                usecase.NewGetUsersUseCase(),
		UpdateUserUseCase:              usecase.NewUpdateUserUseCase(),
		DeleteUserUseCase:              usecase.NewDeleteUserUseCase(),
		UpdateUserPasswordUseCase:      usecase.NewUpdateUserPasswordUseCase(),
		CreateUserValidationUseCase:    usecase.NewCreateUserValidationUseCase(),
		GetUserValidationByHashUseCase: usecase.NewGetUserValidationByHashUseCase(),
		UpdateUserValidationUsed:       usecase.NewUpdateUserValidationUsedUseCase(),
	}
}

func (s *UserService) CreateUser(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error) {
	input := usecase.CreateUserInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.CreateUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	newUserValidation, err := s.CreateUserValidationUseCase.Execute(usecase.CreateUserValidationInputDTO{
		UserID: output.ID,
		Name:   output.Name,
		Email:  output.Email,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	go event.NewUserVerifyEmailEvent(event.UserVerifyEmailEventInputDTO{
		Email: output.Email,
		Name:  output.Name,
		Hash:  newUserValidation.Hash,
	}).Send()

	slog.Info("User created")
	return
}

func (s *UserService) GetUserById(id int32) (output usecase.GetUserOutputDTO, err error) {
	input := usecase.GetUserInputDTO{
		ID: id,
	}

	output, err = s.GetUserUseCase.Execute(input)
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
	output, err = s.GetUserByEmailUseCase.Execute(input)
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

	output, err = s.GetUsersUseCase.Execute(input)
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
	output, err = s.UpdateUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User updated")
	return
}

func (s *UserService) UpdateUserPassword(body io.ReadCloser) (err error) {

	input := RequestUpdateUserPasswordInputDTO{}

	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	userValidation, err := s.GetUserValidationByHashUseCase.Execute(usecase.GetUserValidationByHashInputDTO{
		Hash: input.Hash,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	user, err := s.GetUserUseCase.Execute(usecase.GetUserInputDTO{ID: userValidation.UserID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	_, err = s.UpdateUserPasswordUseCase.Execute(usecase.UpdateUserPasswordInputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: input.Password,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserValidationUsed.Execute(usecase.UpdateUserValidationUsedInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	go event.NewUpdatedPasswordEmailEvent(event.UpdatedPasswordEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
	}).Send()

	slog.Info("User password updated")
	return
}

func (s *UserService) DeleteUser(id int32) (output usecase.DeleteUserOutputDTO, err error) {
	input := usecase.DeleteUserInputDTO{
		ID: id,
	}

	output, err = s.DeleteUserUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User deleted")
	return
}
