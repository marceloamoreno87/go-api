package repository

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AvatarRepository struct {
	db config.SQLCInterface
}

func NewAvatarRepository(db config.SQLCInterface) *AvatarRepository {
	return &AvatarRepository{
		db: db,
	}
}

func (repo *AvatarRepository) CreateAvatar(avatar entityInterface.AvatarInterface) (err error) {
	err = repo.db.GetDbQueries().CreateAvatar(repo.db.GetCtx(), avatar.GetSVG())
	return
}

func (repo *AvatarRepository) GetAvatar(id int32) (output entityInterface.AvatarInterface, err error) {
	a, err := repo.db.GetDbQueries().GetAvatar(repo.db.GetCtx(), id)
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
	a, err := repo.db.GetDbQueries().GetAvatars(repo.db.GetCtx(), db.GetAvatarsParams{
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

func (repo *AvatarRepository) UpdateAvatar(avatar entityInterface.AvatarInterface, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateAvatar(repo.db.GetCtx(), db.UpdateAvatarParams{
		ID:  id,
		Svg: avatar.GetSVG(),
	})
}

func (repo *AvatarRepository) DeleteAvatar(id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeleteAvatar(repo.db.GetCtx(), id)
}
