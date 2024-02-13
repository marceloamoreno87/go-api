package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type UserRepository struct {
	DBConn    *sql.DB
	DBQueries db.Querier
}

func NewUserRepository(DBConn *sql.DB) *UserRepository {
	return &UserRepository{
		DBConn:    DBConn,
		DBQueries: db.New(DBConn),
	}
}

func (repo *UserRepository) CreateUser(user *entity.User) (err error) {
	err = repo.DBQueries.CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) GetUser(id int32) (*entity.User, error) {
	u, err := repo.DBQueries.GetUser(context.Background(), id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	u, err := repo.DBQueries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return &entity.User{}, err
	}
	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (repo *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
	us, err := repo.DBQueries.GetUsers(context.Background(), db.GetUsersParams{
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
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return
}

func (repo *UserRepository) UpdateUser(user *entity.User, id int32) (err error) {
	err = repo.DBQueries.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) DeleteUser(id int32) (err error) {
	err = repo.DBQueries.DeleteUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}
