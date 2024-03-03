package repository

import (
	"context"
	"database/sql"

	"github.com/marceloamoreno/goapi/internal/domain/avatar/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
	"github.com/marceloamoreno/goapi/internal/shared/repository"
)

type AvatarRepositoryInterface interface {
	CreateAvatar(avatar *entity.Avatar) (err error)
	GetAvatar(id int32) (*entity.Avatar, error)
	GetAvatars(limit int32, offset int32) (avatars []*entity.Avatar, err error)
	UpdateAvatar(avatar *entity.Avatar, id int32) (err error)
	DeleteAvatar(id int32) (err error)
	repository.RepositoryInterface
}

type AvatarRepository struct {
	repository.Repository
}

func NewAvatarRepository(dbConn *sql.DB) *AvatarRepository {
	return &AvatarRepository{
		Repository: *repository.NewRepository(dbConn),
	}
}

func (repo *AvatarRepository) CreateAvatar(avatar *entity.Avatar) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).CreateAvatar(context.Background(), avatar.SVG)
	return
}

func (repo *AvatarRepository) GetAvatar(id int32) (avatar *entity.Avatar, err error) {
	p, err := repo.Repository.GetDbQueries().GetAvatar(context.Background(), id)
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

func (repo *AvatarRepository) GetAvatars(limit int32, offset int32) (avatars []*entity.Avatar, err error) {
	ps, err := repo.Repository.GetDbQueries().GetAvatars(context.Background(), db.GetAvatarsParams{
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

func (repo *AvatarRepository) UpdateAvatar(avatar *entity.Avatar, id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).UpdateAvatar(context.Background(), db.UpdateAvatarParams{
		Svg: avatar.SVG,
		ID:  id,
	})
	return
}

func (repo *AvatarRepository) DeleteAvatar(id int32) (err error) {
	err = repo.Repository.GetDbQueries().WithTx(repo.Repository.GetTx()).DeleteAvatar(context.Background(), id)
	return
}
