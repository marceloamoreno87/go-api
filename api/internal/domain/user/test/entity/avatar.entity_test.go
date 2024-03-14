package entity_test

import (
	"testing"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewAvatar(t *testing.T) {
	avatar, err := entity.NewAvatar("svg")
	assert.NoError(t, err)
	assert.NotNil(t, avatar)
}

func TestValidateSVG(t *testing.T) {
	avatar, err := entity.NewAvatar("")
	assert.Error(t, err)
	assert.Nil(t, avatar)
	assert.Equal(t, "[avatar.entity.svg]: SVG is required", err.Error())
}

func TestGetID(t *testing.T) {
	avatar, _ := entity.NewAvatar("svg")
	avatar.SetID(1)
	assert.Equal(t, int32(1), avatar.GetID())
}

func TestGetSVG(t *testing.T) {
	avatar, _ := entity.NewAvatar("svg")
	assert.Equal(t, "svg", avatar.GetSVG())
}
