package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl interface {
	GetLastBook(ctx context.Context) (*domain.Book, error)
	GetListBookByAuthorID(ctx context.Context, authorID int) ([]*domain.Book, error)
	DeleteBookByAuthorID(ctx context.Context, authorID, bookID int) error
	GetByID(ctx context.Context, id int) (*domain.Book, error)
	Delete(ctx context.Context, id int) error
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
	var books *domain.Book

	if err := r.tx(ctx).Order("created_at DESC").First(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) GetListBookByAuthorID(ctx context.Context, authorID int) ([]*domain.Book, error) {
	var books []*domain.Book

	db := r.tx(ctx).Model(&domain.Book{}).Where("author_id = ?", authorID).Find(&books)
	if err := db.Error; err != nil {
		return nil, err
	}

	return books, nil

}

func (r *BookRepository) DeleteBookByAuthorID(ctx context.Context, authorID int, bookID int) error {
	return r.tx(ctx).Model(&domain.Book{}).Delete("id = ? and author_id = ?", bookID, authorID).Error
}

func (r *BookRepository) Delete(ctx context.Context, id int) error {
	return r.tx(ctx).WithContext(ctx).Model(&domain.Book{}).Delete("id = ?", id).Error
}

func (r *BookRepository) GetByID(ctx context.Context, id int) (*domain.Book, error) {
	var book domain.Book
	db := r.tx(ctx).Model(&domain.Book{}).Where("id = ?", id).First(&book)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) Update(ctx context.Context, req *domain.Book) error {
	return r.tx(ctx).Omit("id").Model(&domain.Book{}).Where("id = ?", req.ID).Updates(&req).Error
}

func (r *BookRepository) Create(ctx context.Context, req *domain.Book) error {
	db := r.tx(ctx).Model(&domain.Book{}).Create(&req)

	return db.Error
}
