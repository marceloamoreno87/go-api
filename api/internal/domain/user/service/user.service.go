package service

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

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
	UpdateUserActive               usecaseInterface.UpdateUserActiveUseCaseInterface
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
		UpdateUserActive:               usecase.NewUpdateUserActiveUseCase(),
	}
}

func (s *UserService) CreateUser(input request.RequestCreateUser) (output usecase.CreateUserOutputDTO, err error) {
	output, err = s.CreateUserUseCase.Execute(usecase.CreateUserInputDTO{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
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

func (s *UserService) GetUser(input request.RequestGetUser) (output usecase.GetUserOutputDTO, err error) {
	output, err = s.GetUserUseCase.Execute(usecase.GetUserInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User found")
	return
}

func (s *UserService) GetUsers(input request.RequestGetUsers) (output []usecase.GetUsersOutputDTO, err error) {
	output, err = s.GetUsersUseCase.Execute(usecase.GetUsersInputDTO{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Users found")
	return
}

func (s *UserService) UpdateUser(input request.RequestUpdateUser) (output usecase.UpdateUserOutputDTO, err error) {
	output, err = s.UpdateUserUseCase.Execute(usecase.UpdateUserInputDTO{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User updated")
	return
}

func (s *UserService) UpdateUserPassword(input request.RequestUpdateUserPassword) (err error) {
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

func (s *UserService) DeleteUser(input request.RequestDeleteUser) (output usecase.DeleteUserOutputDTO, err error) {
	output, err = s.DeleteUserUseCase.Execute(usecase.DeleteUserInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User deleted")
	return
}

func (s *UserService) VerifyUser(input request.RequestVerifyUser) (err error) {
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

	err = s.UpdateUserActive.Execute(usecase.UpdateUserActiveInputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   true,
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

	go event.NewUserVerifiedEmailEvent(
		event.UserVerifiedEmailEventInputDTO{
			Email: user.Email,
			Name:  user.Name,
		}).Send()

	slog.Info("User verified")
	return
}

func (s *UserService) ForgotPassword(input request.RequestForgotPassword) (err error) {
	user, err := s.GetUserByEmailUseCase.Execute(usecase.GetUserByEmailInputDTO{Email: input.Email})
	if err != nil {
		slog.Info("err", err)
		return
	}

	userValidation, err := s.CreateUserValidationUseCase.Execute(usecase.CreateUserValidationInputDTO{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	go event.NewPasswordForgotEmailEvent(event.PasswordForgotEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
		Hash:  userValidation.Hash,
	}).Send()

	slog.Info("User forgot password")
	return
}
