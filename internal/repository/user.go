package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository/dto"
	"gorm.io/gorm"
)

type UserRepositoryImpl interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error)
	RegisterUser(ctx context.Context, req *domain.User) error
}

type UserRepository struct {
	TransactionRepository
}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		TransactionRepository: TransactionRepository{
			db: db,
		},
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	var userDTO dto.UserDTO
	db := r.tx(ctx).Table("users").
		Select("id, username, email, password").
		Where("id = ?", id).
		First(&userDTO)
	if err := db.Error; err != nil {
		return nil, err
	}

	return userDTO.TransformToDomain(), nil
}

func (r *UserRepository) GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error) {
	var userDTO dto.UserDTO

	db := r.tx(ctx).Table("users").
		Select("id, username, email, password")

	if req.Username != "" && req.Email != "" {
		db = db.Where("username = ? OR email = ?", req.Username, req.Email)
	} else if req.Username != "" {
		db = db.Where("username = ?", req.Username)
	} else if req.Email != "" {
		db = db.Where("email = ?", req.Email)
	}

	db = db.First(&userDTO)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return userDTO.TransformToDomain(), nil
}

func (r *UserRepository) RegisterUser(ctx context.Context, req *domain.User) error {
	var userDTO dto.UserDTO
	userDTO.TransformFromDomain(req)

	db := r.tx(ctx).Table("users").
		Create(&userDTO)

	return db.Error
}
