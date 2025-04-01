package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl interface {
	Create(ctx context.Context, req *domain.Author) error
	GetByName(ctx context.Context, name string) (*domain.Author, error)
	GetByID(ctx context.Context, id int) (*domain.Author, error)
}

type AuthorRepository struct {
	TransactionRepository
}

func NewAuthorRepository(db *gorm.DB) AuthorRepositoryImpl {
	return &AuthorRepository{
		TransactionRepository: TransactionRepository{
			db: db,
		},
	}
}

func (r *AuthorRepository) Create(ctx context.Context, req *domain.Author) error {
	return r.tx(ctx).Model(&domain.Author{}).Create(&req).Error
}

func (r *AuthorRepository) GetByName(ctx context.Context, name string) (*domain.Author, error) {
	var author domain.Author
	db := r.tx(ctx).Model(&author).Where("name ilike ?", name).First(&author)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) GetByID(ctx context.Context, id int) (*domain.Author, error) {
	var author domain.Author
	db := r.tx(ctx).Model(&domain.Author{}).Where("id = ?", id).First(&author)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &author, nil
}
