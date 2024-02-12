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
		RoleID:   user.RoleID,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		RoleID:    repo.RoleID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
	}, nil
}

func (ur *UserRepository) GetUser(id int32) (*entity.User, error) {

	repo, err := ur.DBQueries.GetUser(context.Background(), id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:        repo.ID,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		RoleID:    repo.RoleID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
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
		RoleID:    repo.RoleID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
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
			RoleID:    u.RoleID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return
}

func (ur *UserRepository) UpdateUser(user *entity.User, id int32) (*entity.User, error) {

	u, err := ur.DBQueries.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	})
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

func (ur *UserRepository) DeleteUser(id int32) (*entity.User, error) {
	u, err := ur.DBQueries.DeleteUser(context.Background(), id)

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
