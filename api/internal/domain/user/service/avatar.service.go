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

func (s *AvatarService) GetAvatar(ctx context.Context, input request.GetAvatarRequest) (output response.GetAvatarResponse, err error) {
	avatar, err := s.GetAvatarUseCase.Execute(ctx, usecase.GetAvatarInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	output = response.GetAvatarResponse{
		ID:  avatar.ID,
		SVG: avatar.SVG,
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(ctx context.Context, input request.GetAvatarsRequest) (output []response.GetAvatarResponse, err error) {
	avatars, err := s.GetAvatarsUseCase.Execute(ctx, usecase.GetAvatarsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	for _, avatar := range avatars {
		output = append(output, response.GetAvatarResponse{
			ID:  avatar.ID,
			SVG: avatar.SVG,
		})
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(ctx context.Context, input request.CreateAvatarRequest) (output response.CreateAvatarResponse, err error) {
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
	output = response.CreateAvatarResponse{
		SVG: created.SVG,
	}
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(ctx context.Context, input request.UpdateAvatarRequest) (output response.UpdateAvatarResponse, err error) {
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
	output = response.UpdateAvatarResponse{
		ID:  updated.ID,
		SVG: updated.SVG,
	}
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(ctx context.Context, input request.DeleteAvatarRequest) (output response.DeleteAvatarResponse, err error) {
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
	output = response.DeleteAvatarResponse{
		ID: deleted.ID,
	}
	slog.Info("Avatar deleted")
	return
}
