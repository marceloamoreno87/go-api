package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
	"github.com/marceloamoreno/goapi/internal/shared/repository"
)

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (err error)
	GetUser(id int32) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsers(limit int32, offset int32) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int32) (err error)
	DeleteUser(id int32) (err error)
	RegisterUser(user *entity.User) (userOutput *entity.User, err error)
	UpdatePasswordUser(user *entity.User, id int32) (err error)
	CreateValidationUser(userValidation *entity.UserValidation) (err error)
	GetValidationUser(id int32) (userValidation *entity.UserValidation, err error)
	GetValidationUserByHash(hash string) (userValidation *entity.UserValidation, err error)
	SetUserValidationUsed(id int32) (err error)
	repository.RepositoryInterface
}

type UserRepository struct {
	repository.Repository
}

func NewUserRepository(DB config.DatabaseInterface) *UserRepository {
	return &UserRepository{
		Repository: *repository.NewRepository(DB),
	}
}

func (repo *UserRepository) CreateUser(user *entity.User) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   user.Active,
		RoleID:   user.RoleID,
		AvatarID: user.AvatarID,
	})
	return
}

func (repo *UserRepository) RegisterUser(user *entity.User) (userOutput *entity.User, err error) {
	newUser, err := repo.GetDbQueries().WithTx(repo.GetTx()).RegisterUser(context.Background(), db.RegisterUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   user.Active,
	})
	if err != nil {
		return
	}
	userOutput = &entity.User{
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

func (repo *UserRepository) GetUser(id int32) (user *entity.User, err error) {
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

func (repo *UserRepository) GetUserByEmail(email string) (user *entity.User, err error) {
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

func (repo *UserRepository) GetUsers(limit int32, offset int32) (users []*entity.User, err error) {
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

func (repo *UserRepository) UpdateUser(user *entity.User, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   user.Active,
		RoleID:   user.RoleID,
		AvatarID: user.AvatarID,
	})
	return
}

func (repo *UserRepository) UpdatePasswordUser(user *entity.User, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdatePasswordUser(context.Background(), db.UpdatePasswordUserParams{
		ID:       id,
		Password: user.Password,
	})
	return
}

func (repo *UserRepository) DeleteUser(id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteUser(context.Background(), id)
	return
}
