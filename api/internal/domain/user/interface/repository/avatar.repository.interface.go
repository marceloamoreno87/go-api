package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AvatarRepositoryInterface interface {
	CreateAvatar(ctx context.Context, avatar entityInterface.AvatarInterface) (err error)
	GetAvatar(ctx context.Context, id int32) (output entityInterface.AvatarInterface, err error)
	GetAvatars(ctx context.Context, limit int32, offset int32) (output []entityInterface.AvatarInterface, err error)
	UpdateAvatar(ctx context.Context, avatar entityInterface.AvatarInterface, id int32) (err error)
	DeleteAvatar(ctx context.Context, id int32) (err error)
}
