package entity_test

import (
	"testing"

	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewAuth(t *testing.T) {
	auth := entity.NewAuth()
	assert.NotNil(t, auth)
}
