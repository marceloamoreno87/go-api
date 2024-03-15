package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarService struct {
	DB config.SQLCInterface
}

func NewAvatarService() *AvatarService {
	return &AvatarService{
		DB: config.Sqcl,
	}
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

func (s *AvatarService) CreateAvatar(body io.ReadCloser) (output usecase.CreateAvatarOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.CreateAvatarInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewCreateAvatarUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(id int32, body io.ReadCloser) (output usecase.UpdateAvatarOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.UpdateAvatarInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewUpdateAvatarUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(id int32) (output usecase.DeleteAvatarOutputDTO, err error) {
	s.DB.Begin()

	input := usecase.DeleteAvatarInputDTO{
		ID: id,
	}

	output, err = usecase.NewDeleteAvatarUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Avatar deleted")
	return
}
