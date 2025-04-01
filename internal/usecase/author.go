package usecase

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
)

type AuthorUseCaseImpl interface {
	CreateAuthorAndBook(ctx context.Context, req *domain.CreateAuthorAndBookRequest) error
	DeleteBookByAuthor(ctx context.Context, id, bookId int) error
	GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error)
	AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error
	CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error
}

type authorUseCase struct {
	config *config.MainConfig
	repo   repository.RepositoryImpl
}

func NewAuthorUseCase(config *config.MainConfig, repo repository.RepositoryImpl) AuthorUseCaseImpl {
	return &authorUseCase{
		config: config,
		repo:   repo,
	}
}

func (u *authorUseCase) CreateAuthorAndBook(ctx context.Context, req *domain.CreateAuthorAndBookRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	return u.repo.GetTransactionRepo().WithTransaction(ctx, func(txCtx context.Context) error {
		err := u.repo.GetAuthorRepo().Create(txCtx, &domain.Author{
			ID:          req.Book.AuthorID,
			Name:        req.Author.Name,
			Email:       req.Author.Email,
			PhoneNumber: req.Author.PhoneNumber,
		})
		if err != nil {
			return err
		}

		err = u.repo.GetBookRepo().Create(txCtx, &domain.Book{
			AuthorID:  req.Book.AuthorID,
			BookName:  req.Book.BookName,
			Title:     req.Book.Title,
			Price:     req.Book.Price,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *authorUseCase) AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	author, err := u.repo.GetAuthorRepo().GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return u.repo.GetBookRepo().Create(ctx, &domain.Book{
		AuthorID: author.ID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (u *authorUseCase) CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	author, err := u.repo.GetAuthorRepo().GetByName(ctx, req.Name)
	if err != nil {
		return err
	}

	if author != nil {
		return errors.New("author already exist")
	}

	return u.repo.GetAuthorRepo().Create(ctx, &domain.Author{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
}

func (u *authorUseCase) DeleteBookByAuthor(ctx context.Context, id int, bookId int) error {
	g, gCtx := errgroup.WithContext(ctx)

	var (
		book   *domain.Book
		author *domain.Author
		err    error
	)

	g.Go(func() error {
		book, err = u.repo.GetBookRepo().GetByID(gCtx, bookId)
		if err != nil {
			return err
		}

		if book == nil {
			return errors.New("book not found")
		}

		return nil
	})

	g.Go(func() error {
		author, err = u.repo.GetAuthorRepo().GetByID(gCtx, id)
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

	return u.repo.GetBookRepo().DeleteBookByAuthorID(ctx, author.ID, book.ID)
}

func (u *authorUseCase) GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error) {
	author, err := u.repo.GetAuthorRepo().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, errors.New("author not found")
	}

	books, err := u.repo.GetBookRepo().GetListBookByAuthorID(ctx, author.ID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
