package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("test", "test@test.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Email, "test@test.com")
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("test", "test@test.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}

func TestUserNameValidate(t *testing.T) {
	user, err := NewUser("", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err.Error(), "Name is required")
}

func TestUserEmailValidate(t *testing.T) {
	user, err := NewUser("test", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err.Error(), "Email is required")
}

func TestUserPasswordValidate(t *testing.T) {
	user, err := NewUser("test", "test@test.com", "")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err.Error(), "Password is required")
}