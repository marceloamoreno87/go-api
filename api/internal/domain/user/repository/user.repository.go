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

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	repo, err := ur.DBQueries.CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		CreatedAt: repo.CreatedAt.Time,
		UpdatedAt: repo.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepository) GetUser(id int64) (*entity.User, error) {
	repo, err := ur.DBQueries.GetUser(context.Background(), id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		CreatedAt: repo.CreatedAt.Time,
		UpdatedAt: repo.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	repo, err := ur.DBQueries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		CreatedAt: repo.CreatedAt.Time,
		UpdatedAt: repo.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
	repo, err := ur.DBQueries.GetUsers(context.Background(), db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []*entity.User{}, err
	}

	for _, u := range repo {
		users = append(users, &entity.User{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt.Time,
			UpdatedAt: u.UpdatedAt.Time,
		})
	}

	return
}

func (ur *UserRepository) UpdateUser(user *entity.User, id int64) (*entity.User, error) {
	repo, err := ur.DBQueries.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		CreatedAt: repo.CreatedAt.Time,
		UpdatedAt: repo.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepository) DeleteUser(id int64) (err error) {
	err = ur.DBQueries.DeleteUser(context.Background(), id)

	if err != nil {
		return
	}
	return
}
