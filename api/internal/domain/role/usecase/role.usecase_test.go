package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/stretchr/testify/mock"
)

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) CreateRole(role *entity.Role) (*entity.Role, error) {
	args := m.Called(role)
	return args.Get(0).(*entity.Role), args.Error(1)
}

func (m *MockRoleRepository) GetRole(id int32) (*entity.Role, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Role), args.Error(1)
}

func (m *MockRoleRepository) GetRoles(limit int32, offset int32) ([]*entity.Role, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*entity.Role), args.Error(1)
}

func (m *MockRoleRepository) UpdateRole(role *entity.Role, id int32) (*entity.Role, error) {
	args := m.Called(role, id)
	return args.Get(0).(*entity.Role), args.Error(1)
}

func (m *MockRoleRepository) DeleteRole(id int32) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoleRepository) GetRoleByInternalName(internal_name string) (*entity.Role, error) {
	args := m.Called(internal_name)
	return args.Get(0).(*entity.Role), args.Error(1)
}
