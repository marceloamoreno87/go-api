package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AvatarRepository struct {
	config.SQLCInterface
}

func NewAvatarRepository() *AvatarRepository {
	return &AvatarRepository{}
}

func (repo *AvatarRepository) CreateAvatar(avatar entityInterface.AvatarInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateAvatar(context.Background(), avatar.GetSVG())
	return
}

func (repo *AvatarRepository) GetAvatar(id int32) (avatar entityInterface.AvatarInterface, err error) {
	p, err := repo.GetDbQueries().GetAvatar(context.Background(), id)
	if err != nil {
		return
	}
	return &entity.Avatar{
		ID:        p.ID,
		SVG:       p.Svg,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (repo *AvatarRepository) GetAvatars(limit int32, offset int32) (avatars []entityInterface.AvatarInterface, err error) {
	ps, err := repo.GetDbQueries().GetAvatars(context.Background(), db.GetAvatarsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return
	}
	for _, p := range ps {
		avatars = append(avatars, &entity.Avatar{
			ID:        p.ID,
			SVG:       p.Svg,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return
}

func (repo *AvatarRepository) UpdateAvatar(avatar entityInterface.AvatarInterface, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateAvatar(context.Background(), db.UpdateAvatarParams{
		Svg: avatar.GetSVG(),
		ID:  id,
	})
	return
}

func (repo *AvatarRepository) DeleteAvatar(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).DeleteAvatar(context.Background(), id)
}
