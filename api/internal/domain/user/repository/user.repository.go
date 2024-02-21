package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/infra/database"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type UserRepository struct {
	database.Repository
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		Repository: *database.NewRepository(dbConn),
	}
}

func (repo *UserRepository) CreateUser(user *entity.User) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
		AvatarID: user.AvatarID,
	})
	return
}

func (repo *UserRepository) GetUser(id int32) (user *entity.User, err error) {
	u, err := repo.Repository.GetDbQueries().GetUser(context.Background(), id)
	if err != nil {
		return
	}

	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (user *entity.User, err error) {
	u, err := repo.Repository.GetDbQueries().GetUserByEmail(context.Background(), email)
	if err != nil {
		return
	}
	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleID:    u.RoleID,
		AvatarID:  u.AvatarID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
	us, err := repo.Repository.GetDbQueries().GetUsers(context.Background(), db.GetUsersParams{
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
			RoleID:    u.RoleID,
			AvatarID:  u.AvatarID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return
}

func (repo *UserRepository) UpdateUser(user *entity.User, id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
		AvatarID: user.AvatarID,
	})
	return
}

func (repo *UserRepository) DeleteUser(id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).DeleteUser(context.Background(), id)
	return
}
