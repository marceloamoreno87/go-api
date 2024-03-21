package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	GetAvatar(ctx context.Context, input request.RequestGetAvatar) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(ctx context.Context, input request.RequestGetAvatars) (output []usecase.GetAvatarsOutputDTO, err error)
	CreateAvatar(ctx context.Context, input request.RequestCreateAvatar) (output usecase.CreateAvatarOutputDTO, err error)
	UpdateAvatar(ctx context.Context, input request.RequestUpdateAvatar) (output usecase.UpdateAvatarOutputDTO, err error)
	DeleteAvatar(ctx context.Context, input request.RequestDeleteAvatar) (output usecase.DeleteAvatarOutputDTO, err error)
}
