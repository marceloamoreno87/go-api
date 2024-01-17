package repository

import (
	"context"

	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	GetUsers(limit int32, offset int32) ([]*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int64) error
}

type UserRepository struct {
	DB *db.Queries
}

func NewUserRepository(DB *db.Queries) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	repo, err := ur.DB.CreateUser(context.Background(), db.CreateUserParams{
		Username: user.Username,
		Password: user.Password,
		Photo:    user.Photo,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Username: repo.Username,
		Password: repo.Password,
		Photo:    repo.Photo,
	}, nil
}

func (ur *UserRepository) GetUser(id int64) (*entity.User, error) {
	repo, err := ur.DB.GetUser(context.Background(), id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Username: repo.Username,
		Password: repo.Password,
		Photo:    repo.Photo,
	}, nil
}

func (ur *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
	repo, err := ur.DB.GetUsers(context.Background(), db.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []*entity.User{}, err
	}

	for _, u := range repo {
		users = append(users, &entity.User{
			ID:       u.ID,
			Username: u.Username,
			Password: u.Password,
			Photo:    u.Photo,
		})
	}

	return
}

func (ur *UserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	repo, err := ur.DB.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Photo:    user.Photo,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Username: repo.Username,
		Password: repo.Password,
		Photo:    repo.Photo,
	}, nil
}

func (ur *UserRepository) DeleteUser(id int64) (err error) {
	err = ur.DB.DeleteUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}
