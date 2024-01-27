package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsers(limit int32, offset int32) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int64) (*entity.User, error)
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
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Name:     repo.Name,
		Email:    repo.Email,
		Password: repo.Password,
	}, nil
}

func (ur *UserRepository) GetUser(id int64) (*entity.User, error) {
	repo, err := ur.DB.GetUser(context.Background(), id)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Name:     repo.Name,
		Email:    repo.Email,
		Password: repo.Password,
	}, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	repo, err := ur.DB.GetUserByEmail(context.Background(), email)
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Name:     repo.Name,
		Email:    repo.Email,
		Password: repo.Password,
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
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		})
	}

	return
}

func (ur *UserRepository) UpdateUser(user *entity.User, id int64) (*entity.User, error) {
	repo, err := ur.DB.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       repo.ID,
		Name:     repo.Name,
		Email:    repo.Email,
		Password: repo.Password,
	}, nil
}

func (ur *UserRepository) DeleteUser(id int64) (err error) {
	err = ur.DB.DeleteUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}
