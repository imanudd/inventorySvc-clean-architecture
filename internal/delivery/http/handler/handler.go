package handler

import (
	"github.com/imanudd/inventorySvc-clean-architecture/internal/usecase"
)

type Handler struct {
	AuthUseCase   usecase.AuthUseCaseImpl
	BookUseCase   usecase.BookUseCaseImpl
	AuthorUseCase usecase.AuthorUseCaseImpl
}

func NewHandler(useCase *Handler) *Handler {
	return &Handler{
		AuthUseCase:   useCase.AuthUseCase,
		BookUseCase:   useCase.BookUseCase,
		AuthorUseCase: useCase.AuthorUseCase,
	}
}
