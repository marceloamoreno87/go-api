package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUser(id int64) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUsers(limit int32, offset int32) ([]*entity.User, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *entity.User, id int64) (*entity.User, error) {
	args := m.Called(user, id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
