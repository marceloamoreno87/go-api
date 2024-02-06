package handler

import (
	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/pkg/api"
)

type PermissionHandler struct {
	HandlerTools         api.HandlerToolsInterface
	PermissionRepository repository.PermissionRepositoryInterface
}

func NewPermissionHandler(PermissionRepository repository.PermissionRepositoryInterface, handlerTools api.HandlerToolsInterface) *PermissionHandler {
	return &PermissionHandler{
		HandlerTools:         handlerTools,
		PermissionRepository: PermissionRepository,
	}
}

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission/{id} [get]

// GetPermission godoc
// @Summary Get Permission
// @Description Get Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission/{id} [get]

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission [get]

// CreatePermission godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param permission body repository.Permission true "Permission"
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission [post]

// UpdatePermission godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param permission body repository.Permission true "Permission"
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission/{id} [put]

// DeletePermission godoc
// @Summary Delete Permission
// @Description Delete Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} api.Response{data=repository.Permission}
// @Failure 400 {object} api.ResponseError
// @Router /permission/{id} [delete]
