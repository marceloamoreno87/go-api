package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserRepository struct {
	config.SQLCInterface
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(user entityInterface.UserInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
	return
}

func (repo *UserRepository) RegisterUser(user entityInterface.UserInterface) (userOutput entityInterface.UserInterface, err error) {
	u, err := repo.GetDbQueries().WithTx(repo.GetTx()).RegisterUser(context.Background(), db.RegisterUserParams{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
	})
	if err != nil {
		return
	}
	userOutput = &entity.User{
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

func (repo *UserRepository) GetUser(id int32) (user entityInterface.UserInterface, err error) {
	u, err := repo.GetDbQueries().GetUser(context.Background(), id)
	if err != nil {
		return
	}

	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Active:    u.Active,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (user entityInterface.UserInterface, err error) {
	u, err := repo.GetDbQueries().GetUserByEmail(context.Background(), email)
	if err != nil {
		return
	}

	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Active:    u.Active,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUsers(limit int32, offset int32) (users []entityInterface.UserInterface, err error) {
	us, err := repo.GetDbQueries().GetUsers(context.Background(), db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, u := range us {
		users = append(users, &entity.User{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			Active:    u.Active,
			RoleID:    u.RoleID,
			AvatarID:  u.AvatarID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return
}

func (repo *UserRepository) UpdateUser(user entityInterface.UserInterface, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
	return
}

func (repo *UserRepository) UpdateUserPassword(user entityInterface.UserInterface, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUserPassword(context.Background(), db.UpdateUserPasswordParams{
		ID:       id,
		Password: user.GetPassword(),
	})
	return
}

func (repo *UserRepository) UpdatedUserValidationUsed(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUserValidationUsed(context.Background(), id)
}

func (repo *UserRepository) DeleteUser(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).DeleteUser(context.Background(), id)
}
