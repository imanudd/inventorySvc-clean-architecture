package handler

import (
	"github.com/imanudd/inventorySvc-clean-architecture/internal/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
