package usecase_test

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUsers() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

