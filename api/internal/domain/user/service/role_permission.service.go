package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionService struct {
	db                                  config.SQLCInterface
	GetRolePermissionsUseCase           usecaseInterface.GetRolePermissionsUseCaseInterface
	CreateRolePermissionUseCase         usecaseInterface.CreateRolePermissionUseCaseInterface
	DeleteRolePermissionByRoleIDUseCase usecaseInterface.DeleteRolePermissionByRoleIDUseCaseInterface
}

func NewRolePermissionService() *RolePermissionService {
	db := config.NewSqlc(config.DB)
	return &RolePermissionService{
		db:                                  db,
		GetRolePermissionsUseCase:           usecase.NewGetRolePermissionsUseCase(db),
		CreateRolePermissionUseCase:         usecase.NewCreateRolePermissionUseCase(db),
		DeleteRolePermissionByRoleIDUseCase: usecase.NewDeleteRolePermissionByRoleIDUseCase(db),
	}
}

func (s *RolePermissionService) GetRolePermissions(ctx context.Context, input request.RequestGetRolePermission) (output []usecase.GetRolePermissionsOutputDTO, err error) {
	output, err = s.GetRolePermissionsUseCase.Execute(ctx, usecase.GetRolePermissionsInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(ctx context.Context, input request.RequestCreateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.CreateRolePermissionUseCase.Execute(ctx, usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	tx.Commit()
	slog.Info("Role permission created")
	return
}

func (s *RolePermissionService) UpdateRolePermission(ctx context.Context, input request.RequestUpdateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	_, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(ctx, usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}

	output, err = s.CreateRolePermissionUseCase.Execute(ctx, usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
	})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	tx.Commit()
	slog.Info("Role permission updated")
	return
}

func (s *RolePermissionService) DeleteRolePermissionByRoleID(ctx context.Context, input request.RequestDeleteRolePermissionByRoleID) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(ctx, usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		tx.Rollback()
		slog.Info("err", err)
		return
	}
	tx.Commit()
	slog.Info("Role permission deleted")
	return
}
