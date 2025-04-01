package repository

import (
	"context"
	"fmt"

	"github.com/imanudd/inventorySvc-clean-architecture/pkg/auth"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl interface {
	WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryImpl {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	txCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx := r.db.WithContext(txCtx).Begin()

	defer func() {
		if recover() != nil || ctx.Done() != nil {
			tx.Rollback()
		}
	}()

	err := fn(auth.SetTrx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error db %v", err)
	}

	return tx.Commit().Error
}

func (r *TransactionRepository) tx(ctx context.Context) *gorm.DB {
	conn := auth.GetTxContext(ctx)
	if conn == nil {
		conn = r.db.WithContext(ctx)
	}

	return conn.WithContext(ctx)
}
