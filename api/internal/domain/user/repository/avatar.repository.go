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

func (repo *AvatarRepository) CreateAvatar(avatar entityInterface.AvatarInterface) (output entityInterface.AvatarInterface, err error) {
	a, err := repo.GetDbQueries().WithTx(repo.GetTx()).CreateAvatar(context.Background(), avatar.GetSVG())
	if err != nil {
		return
	}
	output = &entity.Avatar{
		ID:  a.ID,
		SVG: a.Svg,
	}
	return
}

func (repo *AvatarRepository) GetAvatar(id int32) (output entityInterface.AvatarInterface, err error) {
	a, err := repo.GetDbQueries().GetAvatar(context.Background(), id)
	if err != nil {
		return
	}
	output = &entity.Avatar{
		ID:  a.ID,
		SVG: a.Svg,
	}
	return
}

func (repo *AvatarRepository) GetAvatars(limit int32, offset int32) (output []entityInterface.AvatarInterface, err error) {
	a, err := repo.GetDbQueries().GetAvatars(context.Background(), db.GetAvatarsParams{
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

func (repo *AvatarRepository) UpdateAvatar(avatar entityInterface.AvatarInterface, id int32) (output entityInterface.AvatarInterface, err error) {
	a, err := repo.GetDbQueries().WithTx(repo.GetTx()).UpdateAvatar(context.Background(), db.UpdateAvatarParams{
		Svg: avatar.GetSVG(),
		ID:  id,
	})
	if err != nil {
		return
	}
	output = &entity.Avatar{
		ID:  a.ID,
		SVG: a.Svg,
	}
	return
}

func (repo *AvatarRepository) DeleteAvatar(id int32) (output entityInterface.AvatarInterface, err error) {
	a, err := repo.GetDbQueries().WithTx(repo.GetTx()).DeleteAvatar(context.Background(), id)
	if err != nil {
		return
	}
	output = &entity.Avatar{
		ID:  a.ID,
		SVG: a.Svg,
	}
	return
}