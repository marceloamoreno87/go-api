package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/response"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarService struct {
	db                  config.SQLCInterface
	GetAvatarUseCase    usecase.GetAvatarUseCase
	GetAvatarsUseCase   usecase.GetAvatarsUseCase
	CreateAvatarUseCase usecase.CreateAvatarUseCase
	UpdateAvatarUseCase usecase.UpdateAvatarUseCase
	DeleteAvatarUseCase usecase.DeleteAvatarUseCase
}

func NewAvatarService() *AvatarService {
	db := config.NewSqlc(config.DB)
	return &AvatarService{
		db:                  db,
		GetAvatarUseCase:    *usecase.NewGetAvatarUseCase(db),
		GetAvatarsUseCase:   *usecase.NewGetAvatarsUseCase(db),
		CreateAvatarUseCase: *usecase.NewCreateAvatarUseCase(db),
		UpdateAvatarUseCase: *usecase.NewUpdateAvatarUseCase(db),
		DeleteAvatarUseCase: *usecase.NewDeleteAvatarUseCase(db),
	}
}

func (s *AvatarService) GetAvatar(ctx context.Context, input request.RequestGetAvatar) (output response.ResponseGetAvatar, err error) {
	avatar, err := s.GetAvatarUseCase.Execute(ctx, usecase.GetAvatarInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	output = response.ResponseGetAvatar{
		ID:  avatar.ID,
		SVG: avatar.SVG,
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(ctx context.Context, input request.RequestGetAvatars) (output []response.ResponseGetAvatar, err error) {
	avatars, err := s.GetAvatarsUseCase.Execute(ctx, usecase.GetAvatarsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	for _, avatar := range avatars {
		output = append(output, response.ResponseGetAvatar{
			ID:  avatar.ID,
			SVG: avatar.SVG,
		})
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(ctx context.Context, input request.RequestCreateAvatar) (output response.ResponseCreateAvatar, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	created, err := s.CreateAvatarUseCase.Execute(ctx, usecase.CreateAvatarInputDTO{SVG: input.SVG})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.ResponseCreateAvatar{
		SVG: created.SVG,
	}
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(ctx context.Context, input request.RequestUpdateAvatar) (output response.ResponseUpdateAvatar, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	updated, err := s.UpdateAvatarUseCase.Execute(ctx, usecase.UpdateAvatarInputDTO{ID: input.ID, SVG: input.SVG})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.ResponseUpdateAvatar{
		ID:  updated.ID,
		SVG: updated.SVG,
	}
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(ctx context.Context, input request.RequestDeleteAvatar) (output response.ResponseDeleteAvatar, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	deleted, err := s.DeleteAvatarUseCase.Execute(ctx, usecase.DeleteAvatarInputDTO{ID: input.ID})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.ResponseDeleteAvatar{
		ID: deleted.ID,
	}
	slog.Info("Avatar deleted")
	return
}
