package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarServiceInterface interface {
	CreateAvatar(body io.ReadCloser) (err error)
	GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(limit int32, offset int32) (output []usecase.GetAvatarsOutputDTO, err error)
	UpdateAvatar(id int32, body io.ReadCloser) (err error)
	DeleteAvatar(id int32) (err error)
	config.SQLCInterface
}

type AvatarService struct {
	config.SQLCInterface
}

func NewAvatarService() *AvatarService {
	return &AvatarService{}
}

func (s *AvatarService) GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error) {

	input := usecase.GetAvatarInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetAvatarUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(limit int32, offset int32) (output []usecase.GetAvatarsOutputDTO, err error) {

	input := usecase.GetAvatarsInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	output, err = usecase.NewGetAvatarsUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(body io.ReadCloser) (err error) {
	s.Begin()
	input := usecase.CreateAvatarInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateAvatarUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(id int32, body io.ReadCloser) (err error) {
	s.Begin()
	input := usecase.UpdateAvatarInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	if err = usecase.NewUpdateAvatarUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(id int32) (err error) {
	s.Begin()

	input := usecase.DeleteAvatarInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteAvatarUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("Avatar deleted")
	return
}
