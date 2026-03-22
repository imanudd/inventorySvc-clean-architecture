package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository/dto"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl interface {
	Create(ctx context.Context, req *domain.Author) error
	GetByName(ctx context.Context, name string) (*domain.Author, error)
	GetByID(ctx context.Context, id int64) (*domain.Author, error)
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
	var authorDTO dto.AuthorDTO
	authorDTO.TransformFromDomain(req)

	return r.tx(ctx).Table("authors").Create(&authorDTO).Error
}

func (r *AuthorRepository) GetByName(ctx context.Context, name string) (*domain.Author, error) {
	var authorDTO dto.AuthorDTO
	db := r.tx(ctx).Table("authors").
		Select("id, name, email, phone_number").
		Where("name = ?", name).
		First(&authorDTO)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return authorDTO.TransformToDomain(), nil
}

func (r *AuthorRepository) GetByID(ctx context.Context, id int64) (*domain.Author, error) {
	var authorDTO dto.AuthorDTO
	db := r.tx(ctx).Table("authors").
		Select("id, name, email, phone_number").
		Where("id = ?", id).
		First(&authorDTO)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return authorDTO.TransformToDomain(), nil
}
