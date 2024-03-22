package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserService struct {
	db                             config.SQLCInterface
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
	db := config.NewSqlc(config.DB)
	return &UserService{
		db:                             db,
		GetUserByEmailUseCase:          usecase.NewGetUserByEmailUseCase(db),
		CreateUserUseCase:              usecase.NewCreateUserUseCase(db),
		GetUserUseCase:                 usecase.NewGetUserUseCase(db),
		GetUsersUseCase:                usecase.NewGetUsersUseCase(db),
		UpdateUserUseCase:              usecase.NewUpdateUserUseCase(db),
		DeleteUserUseCase:              usecase.NewDeleteUserUseCase(db),
		UpdateUserPasswordUseCase:      usecase.NewUpdateUserPasswordUseCase(db),
		CreateUserValidationUseCase:    usecase.NewCreateUserValidationUseCase(db),
		GetUserValidationByHashUseCase: usecase.NewGetUserValidationByHashUseCase(db),
		UpdateUserValidationUsed:       usecase.NewUpdateUserValidationUsedUseCase(db),
		UpdateUserActive:               usecase.NewUpdateUserActiveUseCase(db),
	}
}

func (s *UserService) CreateUser(ctx context.Context, input request.RequestCreateUser) (output usecase.CreateUserOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.CreateUserUseCase.Execute(ctx, usecase.CreateUserInputDTO{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	newUserValidation, err := s.CreateUserValidationUseCase.Execute(ctx, usecase.CreateUserValidationInputDTO{
		UserID: output.ID,
		Name:   output.Name,
		Email:  output.Email,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	go event.NewUserVerifyEmailEvent(event.UserVerifyEmailEventInputDTO{
		Email: output.Email,
		Name:  output.Name,
		Hash:  newUserValidation.Hash,
	}).Send()

	tx.Commit()
	slog.Info("User created")
	return
}

func (s *UserService) GetUser(ctx context.Context, input request.RequestGetUser) (output usecase.GetUserOutputDTO, err error) {
	output, err = s.GetUserUseCase.Execute(ctx, usecase.GetUserInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("User found")
	return
}

func (s *UserService) GetUsers(ctx context.Context, input request.RequestGetUsers) (output []usecase.GetUsersOutputDTO, err error) {
	output, err = s.GetUsersUseCase.Execute(ctx, usecase.GetUsersInputDTO{
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

func (s *UserService) UpdateUser(ctx context.Context, input request.RequestUpdateUser) (output usecase.UpdateUserOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.UpdateUserUseCase.Execute(ctx, usecase.UpdateUserInputDTO{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	tx.Commit()
	slog.Info("User updated")
	return
}

func (s *UserService) UpdateUserPassword(ctx context.Context, input request.RequestUpdateUserPassword) (err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	userValidation, err := s.GetUserValidationByHashUseCase.Execute(ctx, usecase.GetUserValidationByHashInputDTO{
		Hash: input.Hash,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	user, err := s.GetUserUseCase.Execute(ctx, usecase.GetUserInputDTO{ID: userValidation.UserID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	_, err = s.UpdateUserPasswordUseCase.Execute(ctx, usecase.UpdateUserPasswordInputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: input.Password,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserValidationUsed.Execute(ctx, usecase.UpdateUserValidationUsedInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	go event.NewUpdatedPasswordEmailEvent(event.UpdatedPasswordEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
	}).Send()

	tx.Commit()
	slog.Info("User password updated")
	return
}

func (s *UserService) DeleteUser(ctx context.Context, input request.RequestDeleteUser) (output usecase.DeleteUserOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.DeleteUserUseCase.Execute(ctx, usecase.DeleteUserInputDTO{ID: input.ID})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	tx.Commit()
	slog.Info("User deleted")
	return
}

func (s *UserService) VerifyUser(ctx context.Context, input request.RequestVerifyUser) (err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	userValidation, err := s.GetUserValidationByHashUseCase.Execute(ctx, usecase.GetUserValidationByHashInputDTO{
		Hash: input.Hash,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	user, err := s.GetUserUseCase.Execute(ctx, usecase.GetUserInputDTO{ID: userValidation.UserID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserActive.Execute(ctx, usecase.UpdateUserActiveInputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   true,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserValidationUsed.Execute(ctx, usecase.UpdateUserValidationUsedInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	go event.NewUserVerifiedEmailEvent(
		event.UserVerifiedEmailEventInputDTO{
			Email: user.Email,
			Name:  user.Name,
		}).Send()

	tx.Commit()
	slog.Info("User verified")
	return
}

func (s *UserService) ForgotPassword(ctx context.Context, input request.RequestForgotPassword) (err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	user, err := s.GetUserByEmailUseCase.Execute(ctx, usecase.GetUserByEmailInputDTO{Email: input.Email})
	if err != nil {
		slog.Info("err", err)
		return
	}

	userValidation, err := s.CreateUserValidationUseCase.Execute(ctx, usecase.CreateUserValidationInputDTO{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	go event.NewPasswordForgotEmailEvent(event.PasswordForgotEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
		Hash:  userValidation.Hash,
	}).Send()

	tx.Commit()
	slog.Info("User forgot password")
	return
}
