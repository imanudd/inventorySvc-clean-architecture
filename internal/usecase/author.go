package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
)

type AuthorUseCaseImpl interface {
	DeleteBookByAuthor(ctx context.Context, id, bookId int) error
	GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error)
	AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error
	CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error
}

type authorUseCase struct {
	config     *config.MainConfig
	authorRepo repository.AuthorRepositoryImpl
	bookRepo   repository.BookRepositoryImpl
}

func NewAuthorUseCase(config *config.MainConfig, authorRepo repository.AuthorRepositoryImpl, bookRepo repository.BookRepositoryImpl) AuthorUseCaseImpl {
	return &authorUseCase{
		config:     config,
		authorRepo: authorRepo,
		bookRepo:   bookRepo,
	}
}

func (a *authorUseCase) AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error {
	author, err := a.authorRepo.GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return a.bookRepo.Create(ctx, &domain.Book{
		AuthorID: author.ID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (a *authorUseCase) CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error {
	author, err := a.authorRepo.GetByName(ctx, req.Name)
	if err != nil {
		return err
	}

	if author != nil {
		return errors.New("author already exist")
	}

	return a.authorRepo.Create(ctx, &domain.Author{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
}

func (a *authorUseCase) DeleteBookByAuthor(ctx context.Context, id int, bookId int) error {
	book, err := a.bookRepo.GetByID(ctx, bookId)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	author, err := a.authorRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return a.bookRepo.DeleteBookByAuthorID(ctx, author.ID, book.ID)
}

func (a *authorUseCase) GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error) {
	author, err := a.authorRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	fmt.Println("author id", author)

	if author == nil {
		return nil, errors.New("author not found")
	}

	books, err := a.bookRepo.GetListBookByAuthorID(ctx, author.ID)
	if err != nil {
		return nil, err
	}

	if len(books) < 1 {
		return nil, errors.New("book not found")
	}

	return books, nil
}
