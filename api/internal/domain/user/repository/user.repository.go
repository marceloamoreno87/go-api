package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type Userrepository interface {
	CreateUser(ctx context.Context, user *entity.User) (output *entity.User, err error)
	GetUser(ctx context.Context, id int32) (output *entity.User, err error)
	GetUserByEmail(ctx context.Context, email string) (output *entity.User, err error)
	GetUsers(ctx context.Context, limit int32, offset int32) (output []*entity.User, err error)
	UpdateUser(ctx context.Context, user *entity.User, id int32) (ouerr error)
	DeleteUser(ctx context.Context, id int32) (err error)
	UpdateUserPassword(ctx context.Context, id int32, password string) (err error)
	UpdateUserActive(ctx context.Context, id int32, active bool) (err error)
}

type UserRepository struct {
	db config.SQLCInterface
}

func NewUserRepository(db config.SQLCInterface) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *entity.User) (output *entity.User, err error) {
	newUser, err := repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateUser(ctx, db.CreateUserParams{
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

func (repo *UserRepository) GetUser(ctx context.Context, id int32) (output *entity.User, err error) {
	u, err := repo.db.GetDbQueries().GetUser(ctx, id)
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

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (output *entity.User, err error) {
	u, err := repo.db.GetDbQueries().GetUserByEmail(ctx, email)
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

func (repo *UserRepository) GetUsers(ctx context.Context, limit int32, offset int32) (output []*entity.User, err error) {
	u, err := repo.db.GetDbQueries().GetUsers(ctx, db.GetUsersParams{
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

func (repo *UserRepository) UpdateUser(ctx context.Context, user *entity.User, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateUser(ctx, db.UpdateUserParams{
		ID:       id,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	})
}

func (repo *UserRepository) UpdateUserPassword(ctx context.Context, id int32, password string) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		ID:       id,
		Password: password,
	})
}

func (repo *UserRepository) UpdateUserActive(ctx context.Context, id int32, active bool) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateUserActive(ctx, db.UpdateUserActiveParams{
		ID:     id,
		Active: active,
	})
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeleteUser(ctx, id)
}
