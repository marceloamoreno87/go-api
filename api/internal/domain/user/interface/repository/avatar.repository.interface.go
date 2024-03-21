package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AvatarRepositoryInterface interface {
	CreateAvatar(ctx context.Context, avatar entityInterface.AvatarInterface) (err error)
	GetAvatar(id int32) (output entityInterface.AvatarInterface, err error)
	GetAvatars(limit int32, offset int32) (output []entityInterface.AvatarInterface, err error)
	UpdateAvatar(avatar entityInterface.AvatarInterface, id int32) (err error)
	DeleteAvatar(id int32) (err error)
}
