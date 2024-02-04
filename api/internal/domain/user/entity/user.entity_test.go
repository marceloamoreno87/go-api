package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Test User", "test@example.com", "password")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.GetName())
	assert.Equal(t, "test@example.com", user.GetEmail())
	assert.True(t, user.ComparePassword("password"))
}

func TestValidate(t *testing.T) {
	user := &User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	err := user.Validate()
	assert.NoError(t, err)

	user.Name = ""
	err = user.Validate()
	assert.EqualError(t, err, "Name is required")

	user.Name = "Test User"
	user.Email = ""
	err = user.Validate()
	assert.EqualError(t, err, "Email is required")

	user.Email = "test@example.com"
	user.Password = ""
	err = user.Validate()
	assert.EqualError(t, err, "Password is required")
}

func TestIsEmailValid(t *testing.T) {
	user := &User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	valid, err := user.IsEmailValid()
	assert.NoError(t, err)
	assert.True(t, valid)

	user.Email = "invalid email"
	valid, err = user.IsEmailValid()
	assert.EqualError(t, err, "Email is invalid")
	assert.False(t, valid)
}

func TestComparePassword(t *testing.T) {
	user, _ := NewUser("Test User", "test@example.com", "password")
	assert.True(t, user.ComparePassword("password"))
	assert.False(t, user.ComparePassword("wrongpassword"))
}
