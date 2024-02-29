package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marceloamoreno/goapi/internal/domain/permission/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type PermissionHandler struct {
	response.Responses
	service service.PermissionServiceInterface
}

func NewPermissionHandler(
	service service.PermissionServiceInterface,
) *PermissionHandler {
	return &PermissionHandler{
		service: service,
	}
}

// GetPermission godoc
// @Summary Get Permission
// @Description Get Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} response.Response{data=usecase.GetPermissionOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission/{id} [get]
// @Security     JWT
func (h *PermissionHandler) GetPermission(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	output, err := h.service.GetPermission(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Permission found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))
}

// GetPermissions godoc
// @Summary Get Permissions
// @Description Get Permissions
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetPermissionsOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission [get]
// @Security     JWT
func (h *PermissionHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")

	output, err := h.service.GetPermissions(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Permissions found")
	h.SendResponse(w, h.NewResponse(output, http.StatusOK))

}

// CreateRole godoc
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param role body usecase.CreatePermissionInputDTO true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [post]
// @Security     JWT
func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {

	err := h.service.CreatePermission(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}
	slog.Info("Permission created")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))

}

// UpdateRole godoc
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Param role body usecase.UpdatePermissionInputDTO true "Permission"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /permission/{id} [put]
// @Security     JWT
func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.UpdatePermission(id, r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Permission updated")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))
}

// DeletePermission godoc
// @Summary Delete Permission
// @Description Delete Permission
// @Tags Permission
// @Accept  json
// @Produce  json
// @Param id path string true "Permission ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /permission/{id} [delete]
// @Security     JWT
func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.service.DeletePermission(id)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	slog.Info("Permission deleted")
	h.SendResponse(w, h.NewResponse(nil, http.StatusOK))

}
