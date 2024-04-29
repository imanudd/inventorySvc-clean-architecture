package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
)

type BookUseCaseImpl interface {
	GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, req *domain.UpdateBookRequest) error
	AddBook(ctx context.Context, req *domain.CreateBookRequest) error
}

type bookUseCase struct {
	config     *config.MainConfig
	bookRepo   repository.BookRepositoryImpl
	authorRepo repository.AuthorRepositoryImpl
}

func NewBookUseCase(config *config.MainConfig, bookRepo repository.BookRepositoryImpl, authorRepo repository.AuthorRepositoryImpl) BookUseCaseImpl {
	return &bookUseCase{
		config:     config,
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

func (s *bookUseCase) DeleteBook(ctx context.Context, id int) error {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	return s.bookRepo.Delete(ctx, id)
}

func (s *bookUseCase) GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	author, err := s.authorRepo.GetByID(ctx, book.AuthorID)
	if err != nil {
		return nil, err
	}

	resp := &domain.DetailBook{
		ID:         book.ID,
		AuthorID:   book.AuthorID,
		AuthorName: author.Name,
		BookName:   book.BookName,
		Title:      book.Title,
		Price:      book.Price,
		CreatedAt:  book.CreatedAt,
	}

	return resp, nil
}

func (s *bookUseCase) UpdateBook(ctx context.Context, req *domain.UpdateBookRequest) error {
	book, err := s.bookRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	author, err := s.bookRepo.GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return s.bookRepo.Update(ctx, &domain.Book{
		ID:       req.ID,
		AuthorID: req.AuthorID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (s *bookUseCase) AddBook(ctx context.Context, req *domain.CreateBookRequest) error {
	author, err := s.authorRepo.GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	book := &domain.Book{
		AuthorID:  req.AuthorID,
		BookName:  req.BookName,
		Title:     req.Title,
		Price:     req.Price,
		CreatedAt: time.Now(),
	}

	return s.bookRepo.Create(ctx, book)
}
