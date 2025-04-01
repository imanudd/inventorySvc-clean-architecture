package repository

import (
	"gorm.io/gorm"
)

type RepositoryImpl interface {
	GetUserRepo() UserRepositoryImpl
	GetBookRepo() BookRepositoryImpl
	GetAuthorRepo() AuthorRepositoryImpl
	GetTransactionRepo() TransactionRepositoryImpl
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(database *gorm.DB) RepositoryImpl {
	return &Repository{
		db: database,
	}
}

func (r *Repository) GetUserRepo() UserRepositoryImpl {
	return NewUserRepository(r.db)
}

func (r *Repository) GetBookRepo() BookRepositoryImpl {
	return NewBookRepository(r.db)
}

func (r *Repository) GetAuthorRepo() AuthorRepositoryImpl {
	return NewAuthorRepository(r.db)
}

func (r *Repository) GetTransactionRepo() TransactionRepositoryImpl {
	return NewTransactionRepository(r.db)
}
