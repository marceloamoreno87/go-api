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
}
