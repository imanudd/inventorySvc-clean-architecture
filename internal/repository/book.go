package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository/dto"
	"gorm.io/gorm"
)

type BookRepositoryImpl interface {
	GetLastBook(ctx context.Context) (*domain.Book, error)
	GetListBookByAuthorID(ctx context.Context, authorID int64) ([]*domain.Book, error)
	DeleteBookByAuthorID(ctx context.Context, authorID, bookID int64) error
	GetByID(ctx context.Context, id int64) (*domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req *domain.Book) error
	Create(ctx context.Context, req *domain.Book) error
}

type BookRepository struct {
	TransactionRepository
}

func NewBookRepository(db *gorm.DB) BookRepositoryImpl {
	return &BookRepository{
		TransactionRepository: TransactionRepository{
			db: db,
		},
	}
}

func (r *BookRepository) GetLastBook(ctx context.Context) (*domain.Book, error) {
	var bookDTO dto.BookDTO

	if err := r.tx(ctx).Table("books").
		Select("id, author_id, book_name, title, price, created_at").
		Order("created_at DESC").
		First(&bookDTO).Error; err != nil {
		return nil, err
	}

	return bookDTO.TransformToDomain(), nil
}

func (r *BookRepository) GetListBookByAuthorID(ctx context.Context, authorID int64) ([]*domain.Book, error) {
	var bookDTOs []dto.BookDTO

	db := r.tx(ctx).Table("books").
		Select("id, author_id, book_name, title, price, created_at").
		Where("author_id = ?", authorID).
		Find(&bookDTOs)
	if err := db.Error; err != nil {
		return nil, err
	}

	books := make([]*domain.Book, len(bookDTOs))
	for i, bookDTO := range bookDTOs {
		books[i] = bookDTO.TransformToDomain()
	}

	return books, nil
}

func (r *BookRepository) DeleteBookByAuthorID(ctx context.Context, authorID int64, bookID int64) error {
	return r.tx(ctx).Table("books").Delete("id = ? and author_id = ?", bookID, authorID).Error
}

func (r *BookRepository) Delete(ctx context.Context, id int64) error {
	return r.tx(ctx).WithContext(ctx).Table("books").Delete("id = ?", id).Error
}

func (r *BookRepository) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	var bookDTO dto.BookDTO
	db := r.tx(ctx).Table("books").
		Select("id, author_id, book_name, title, price, created_at").
		Where("id = ?", id).
		First(&bookDTO)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return bookDTO.TransformToDomain(), nil
}

func (r *BookRepository) Update(ctx context.Context, req *domain.Book) error {
	var bookDTO dto.BookDTO
	bookDTO.TransformFromDomain(req)

	return r.tx(ctx).Table("books").
		Where("id = ?", req.ID).
		Updates(&bookDTO).Error
}

func (r *BookRepository) Create(ctx context.Context, req *domain.Book) error {
	var bookDTO dto.BookDTO
	bookDTO.TransformFromDomain(req)

	db := r.tx(ctx).Table("books").Create(&bookDTO)

	return db.Error
}
