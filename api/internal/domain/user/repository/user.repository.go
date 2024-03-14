package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserRepository struct {
	config.SQLCInterface
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(user entityInterface.UserInterface) (output db.User, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) RegisterUser(user entityInterface.UserInterface) (output db.User, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).RegisterUser(context.Background(), db.RegisterUserParams{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUser(id int32) (output db.User, err error) {
	output, err = repo.GetDbQueries().GetUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUserByEmail(email string) (output db.User, err error) {
	output, err = repo.GetDbQueries().GetUserByEmail(context.Background(), email)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUsers(limit int32, offset int32) (output []db.User, err error) {
	output, err = repo.GetDbQueries().GetUsers(context.Background(), db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) UpdateUser(user entityInterface.UserInterface, id int32) (output db.User, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) UpdateUserPassword(user entityInterface.UserInterface, id int32) (output db.User, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUserPassword(context.Background(), db.UpdateUserPasswordParams{
		ID:       id,
		Password: user.GetPassword(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) DeleteUser(id int32) (output db.User, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}
