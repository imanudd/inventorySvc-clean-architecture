package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
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
	var user domain.User
	db := r.tx(ctx).Model(&user).Where("id = ?", id).First(&user)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error) {
	var user *domain.User
	db := r.tx(ctx).Model(&domain.User{}).Where("username ilike ? or email ilike ?", req.Username, req.Email).First(&user)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) RegisterUser(ctx context.Context, req *domain.User) error {
	db := r.tx(ctx).Model(&domain.User{}).Create(&req)

	return db.Error
}
