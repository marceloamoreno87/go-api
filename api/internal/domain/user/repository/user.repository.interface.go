package repository

import (
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
)

type UserRepositoryInterface interface {
	SetTx(tx *sql.Tx)
	CreateUser(user *entity.User) (err error)
	GetUser(id int32) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsers(limit int32, offset int32) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int32) (err error)
	DeleteUser(id int32) (err error)
}
