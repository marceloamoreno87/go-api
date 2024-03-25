package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/response"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserService struct {
	db                             config.SQLCInterface
	GetUserByEmailUseCase          usecase.GetUserByEmailUseCase
	CreateUserUseCase              usecase.CreateUserUseCase
	GetUserUseCase                 usecase.GetUserUseCase
	GetUsersUseCase                usecase.GetUsersUseCase
	UpdateUserUseCase              usecase.UpdateUserUseCase
	DeleteUserUseCase              usecase.DeleteUserUseCase
	UpdateUserPasswordUseCase      usecase.UpdateUserPasswordUseCase
	CreateUserValidationUseCase    usecase.CreateUserValidationUseCase
	GetUserValidationByHashUseCase usecase.GetUserValidationByHashUseCase
	UpdateUserValidationUsed       usecase.UpdateUserValidationUsedUseCase
	UpdateUserActive               usecase.UpdateUserActiveUseCase
}

func NewUserService() *UserService {
	db := config.NewSqlc(config.DB)
	return &UserService{
		db:                             db,
		GetUserByEmailUseCase:          *usecase.NewGetUserByEmailUseCase(db),
		CreateUserUseCase:              *usecase.NewCreateUserUseCase(db),
		GetUserUseCase:                 *usecase.NewGetUserUseCase(db),
		GetUsersUseCase:                *usecase.NewGetUsersUseCase(db),
		UpdateUserUseCase:              *usecase.NewUpdateUserUseCase(db),
		DeleteUserUseCase:              *usecase.NewDeleteUserUseCase(db),
		UpdateUserPasswordUseCase:      *usecase.NewUpdateUserPasswordUseCase(db),
		CreateUserValidationUseCase:    *usecase.NewCreateUserValidationUseCase(db),
		GetUserValidationByHashUseCase: *usecase.NewGetUserValidationByHashUseCase(db),
		UpdateUserValidationUsed:       *usecase.NewUpdateUserValidationUsedUseCase(db),
		UpdateUserActive:               *usecase.NewUpdateUserActiveUseCase(db),
	}
}

func (s *UserService) CreateUser(ctx context.Context, input request.CreateUserRequest) (output response.CreateUserResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	created, err := s.CreateUserUseCase.Execute(ctx, usecase.CreateUserInputDTO{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	newUserValidation, err := s.CreateUserValidationUseCase.Execute(ctx, usecase.CreateUserValidationInputDTO{
		UserID: output.ID,
		Name:   output.Name,
		Email:  output.Email,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	go event.NewUserVerifyEmailEvent(event.UserVerifyEmailEventInputDTO{
		Email: output.Email,
		Name:  output.Name,
		Hash:  newUserValidation.Hash,
	}).Send()

	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.CreateUserResponse{
		ID:       created.ID,
		Name:     created.Name,
		Email:    created.Email,
		AvatarID: created.AvatarID,
		RoleID:   created.RoleID,
	}
	slog.Info("User created")
	return
}

func (s *UserService) GetUser(ctx context.Context, input request.GetUserRequest) (output response.GetUserResponse, err error) {
	user, err := s.GetUserUseCase.Execute(ctx, usecase.GetUserInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	output = response.GetUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		AvatarID: user.AvatarID,
		RoleID:   user.RoleID,
	}
	slog.Info("User found")
	return
}

func (s *UserService) GetUsers(ctx context.Context, input request.GetUsersRequest) (output []response.GetUsersResponse, err error) {
	users, err := s.GetUsersUseCase.Execute(ctx, usecase.GetUsersInputDTO{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	for _, user := range users {
		output = append(output, response.GetUsersResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			AvatarID: user.AvatarID,
			RoleID:   user.RoleID,
		})
	}
	slog.Info("Users found")
	return
}

func (s *UserService) UpdateUser(ctx context.Context, input request.UpdateUserRequest) (output response.UpdateUserResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	updated, err := s.UpdateUserUseCase.Execute(ctx, usecase.UpdateUserInputDTO{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.UpdateUserResponse{
		ID:       updated.ID,
		Name:     updated.Name,
		Email:    updated.Email,
		AvatarID: updated.AvatarID,
		RoleID:   updated.RoleID,
	}
	slog.Info("User updated")
	return
}

func (s *UserService) UpdateUserPassword(ctx context.Context, input request.UpdateUserPasswordRequest) (err error) {
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
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserValidationUsed.Execute(ctx, usecase.UpdateUserValidationUsedInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	go event.NewUpdatedPasswordEmailEvent(event.UpdatedPasswordEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
	}).Send()

	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("User password updated")
	return
}

func (s *UserService) DeleteUser(ctx context.Context, input request.DeleteUserRequest) (output response.DeleteUserResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	deleted, err := s.DeleteUserUseCase.Execute(ctx, usecase.DeleteUserInputDTO{ID: input.ID})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.DeleteUserResponse{
		ID: deleted.ID,
	}
	slog.Info("User deleted")
	return
}

func (s *UserService) VerifyUser(ctx context.Context, input request.VerifyUserRequest) (err error) {
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
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	err = s.UpdateUserValidationUsed.Execute(ctx, usecase.UpdateUserValidationUsedInputDTO{
		UserID: user.ID,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	go event.NewUserVerifiedEmailEvent(
		event.UserVerifiedEmailEventInputDTO{
			Email: user.Email,
			Name:  user.Name,
		}).Send()

	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("User verified")
	return
}

func (s *UserService) ForgotPassword(ctx context.Context, input request.ForgotPasswordRequest) (err error) {
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
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	go event.NewPasswordForgotEmailEvent(event.PasswordForgotEmailEventInputDTO{
		Email: user.Email,
		Name:  user.Name,
		Hash:  userValidation.Hash,
	}).Send()

	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("User forgot password")
	return
}
