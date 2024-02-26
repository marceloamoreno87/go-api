package handler

type HandlerInterface interface {
}

type Handler struct {
	Repository interface{}
	Service    interface{}
}

func NewHandler() *Handler {
	return &Handler{
		Repository: nil,
		Service:    nil,
	}
}

func (h *Handler) SetRepository(repository interface{}) {
	h.Repository = repository
}

func (h *Handler) SetService(service interface{}) {
	h.Service = service
}
