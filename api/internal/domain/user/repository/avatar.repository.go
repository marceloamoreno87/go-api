package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type Avatarrepository interface {
	CreateAvatar(ctx context.Context, avatar *entity.Avatar) (err error)
	GetAvatar(ctx context.Context, id int32) (output *entity.Avatar, err error)
	GetAvatars(ctx context.Context, limit int32, offset int32) (output []*entity.Avatar, err error)
	UpdateAvatar(ctx context.Context, avatar *entity.Avatar, id int32) (err error)
	DeleteAvatar(ctx context.Context, id int32) (err error)
}

type AvatarRepository struct {
	db config.SQLCInterface
}

func NewAvatarRepository(db config.SQLCInterface) *AvatarRepository {
	return &AvatarRepository{
		db: db,
	}
}

func (repo *AvatarRepository) CreateAvatar(ctx context.Context, avatar *entity.Avatar) (err error) {
	err = repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateAvatar(ctx, avatar.GetSVG())
	return
}

func (repo *AvatarRepository) GetAvatar(ctx context.Context, id int32) (output *entity.Avatar, err error) {
	a, err := repo.db.GetDbQueries().GetAvatar(ctx, id)
	if err != nil {
		return
	}
	output = &entity.Avatar{
		ID:  a.ID,
		SVG: a.Svg,
	}
	return
}

func (repo *AvatarRepository) GetAvatars(ctx context.Context, limit int32, offset int32) (output []*entity.Avatar, err error) {
	a, err := repo.db.GetDbQueries().GetAvatars(ctx, db.GetAvatarsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, avatar := range a {
		output = append(output, &entity.Avatar{
			ID:  avatar.ID,
			SVG: avatar.Svg,
		})
	}
	return
}

func (repo *AvatarRepository) UpdateAvatar(ctx context.Context, avatar *entity.Avatar, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateAvatar(ctx, db.UpdateAvatarParams{
		ID:  id,
		Svg: avatar.GetSVG(),
	})
}

func (repo *AvatarRepository) DeleteAvatar(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeleteAvatar(ctx, id)
}
