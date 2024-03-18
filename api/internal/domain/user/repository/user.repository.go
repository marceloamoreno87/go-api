package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserRepository struct {
	DB config.SQLCInterface
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: config.NewSqlc(config.DB),
	}
}

func (repo *UserRepository) CreateUser(user entityInterface.UserInterface) (output entityInterface.UserInterface, err error) {
	newUser, err := repo.DB.GetDbQueries().CreateUser(context.Background(), db.CreateUserParams{
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
	output = &entity.User{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Active:    newUser.Active,
		RoleID:    newUser.RoleID,
		AvatarID:  newUser.AvatarID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	return
}

func (repo *UserRepository) GetUser(id int32) (output entityInterface.UserInterface, err error) {
	u, err := repo.DB.GetDbQueries().GetUser(context.Background(), id)
	if err != nil {
		return
	}
	output = &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Active:    u.Active,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return
}

func (repo *UserRepository) GetUserByEmail(email string) (output entityInterface.UserInterface, err error) {
	u, err := repo.DB.GetDbQueries().GetUserByEmail(context.Background(), email)
	if err != nil {
		return
	}
	output = &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Active:    u.Active,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return
}

func (repo *UserRepository) GetUsers(limit int32, offset int32) (output []entityInterface.UserInterface, err error) {
	u, err := repo.DB.GetDbQueries().GetUsers(context.Background(), db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, user := range u {
		output = append(output, &entity.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Active:    user.Active,
			RoleID:    user.RoleID,
			AvatarID:  user.AvatarID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return
}

func (repo *UserRepository) UpdateUser(user entityInterface.UserInterface, id int32) (err error) {
	return repo.DB.GetDbQueries().UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
}

func (repo *UserRepository) UpdateUserPassword(id int32, password string) (err error) {
	return repo.DB.GetDbQueries().UpdateUserPassword(context.Background(), db.UpdateUserPasswordParams{
		ID:       id,
		Password: password,
	})
}

func (repo *UserRepository) UpdateUserActive(id int32, active bool) (err error) {
	return repo.DB.GetDbQueries().UpdateUserActive(context.Background(), db.UpdateUserActiveParams{
		ID:     id,
		Active: active,
	})
}

func (repo *UserRepository) DeleteUser(id int32) (err error) {
	return repo.DB.GetDbQueries().DeleteUser(context.Background(), id)
}
