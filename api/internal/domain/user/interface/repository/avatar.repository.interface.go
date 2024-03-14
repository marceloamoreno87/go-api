package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AvatarRepositoryInterface interface {
	CreateAvatar(avatar entityInterface.AvatarInterface) (err error)
	GetAvatar(id int32) (entityInterface.AvatarInterface, error)
	GetAvatars(limit int32, offset int32) (avatars []entityInterface.AvatarInterface, err error)
	UpdateAvatar(avatar entityInterface.AvatarInterface, id int32) (err error)
	DeleteAvatar(id int32) (err error)
	config.SQLCInterface
}
