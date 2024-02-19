package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/pkg/sqlc/db"
)

type UserRepository struct {
	dbConn    *sql.DB
	dbQueries *db.Queries
	tx        *sql.Tx
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		dbConn:    dbConn,
		dbQueries: db.New(dbConn),
		tx:        nil,
	}
}

func (ur *UserRepository) SetTx(tx *sql.Tx) {
	ur.tx = tx
}

func (ur *UserRepository) CreateUser(user *entity.User) (err error) {
	err = ur.dbQueries.WithTx(ur.tx).CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	})
	return
}

func (ur *UserRepository) GetUser(id int32) (*entity.User, error) {
	u, err := ur.dbQueries.GetUser(context.Background(), id)
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

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	u, err := ur.dbQueries.GetUserByEmail(context.Background(), email)
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

func (ur *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
	us, err := ur.dbQueries.GetUsers(context.Background(), db.GetUsersParams{
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

func (ur *UserRepository) UpdateUser(user *entity.User, id int32) (err error) {
	err = ur.dbQueries.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	})
	return
}

func (ur *UserRepository) DeleteUser(id int32) (err error) {
	err = ur.dbQueries.DeleteUser(context.Background(), id)
	return
}
