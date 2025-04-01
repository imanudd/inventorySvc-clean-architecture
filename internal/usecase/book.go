package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
	"golang.org/x/sync/errgroup"
)

type BookUseCaseImpl interface {
	GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, req *domain.UpdateBookRequest) error
	AddBook(ctx context.Context, req *domain.CreateBookRequest) error
}

type bookUseCase struct {
	config *config.MainConfig
	repo   repository.RepositoryImpl
}

func NewBookUseCase(config *config.MainConfig, repo repository.RepositoryImpl) BookUseCaseImpl {
	return &bookUseCase{
		config: config,
		repo:   repo,
	}
}

func (s *bookUseCase) DeleteBook(ctx context.Context, id int) error {

	book, err := s.repo.GetBookRepo().GetByID(ctx, id)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	return s.repo.GetBookRepo().Delete(ctx, id)
}

func (s *bookUseCase) GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error) {
	book, err := s.repo.GetBookRepo().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	author, err := s.repo.GetAuthorRepo().GetByID(ctx, book.AuthorID)
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
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	g, gCtx := errgroup.WithContext(ctx)

	var (
		book   *domain.Book
		author *domain.Author
		err    error
	)

	g.Go(func() error {
		book, err = s.repo.GetBookRepo().GetByID(gCtx, req.ID)
		if err != nil {
			return err
		}

		if book == nil {
			return errors.New("book not found")
		}
		return nil
	})

	g.Go(func() error {
		author, err = s.repo.GetAuthorRepo().GetByID(gCtx, req.AuthorID)
		if err != nil {
			return err
		}

		if author == nil {
			return errors.New("author not found")
		}

		return nil
	})

	if err = g.Wait(); err != nil {
		return err
	}

	return s.repo.GetBookRepo().Update(ctx, &domain.Book{
		ID:       req.ID,
		AuthorID: req.AuthorID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (s *bookUseCase) AddBook(ctx context.Context, req *domain.CreateBookRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	author, err := s.repo.GetAuthorRepo().GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	// books, err := s.repo.BookRepository.GetLastBook(ctx)
	// if err != nil {
	// 	return err
	// }

	// err = s.es.Save(ctx, elasticsearch.BOOK_DETAILS, &domain.CreateDetailBook{
	// 	Id:       books.ID + 1,
	// 	BookName: req.BookName,
	// 	Title:    req.Title,
	// 	Price:    req.Price,
	// 	Author:   *author,
	// })

	if err != nil {
		return err
	}

	book := &domain.Book{
		AuthorID:  req.AuthorID,
		BookName:  req.BookName,
		Title:     req.Title,
		Price:     req.Price,
		CreatedAt: time.Now(),
	}

	return s.repo.GetBookRepo().Create(ctx, book)
}
