package dto

import (
	"database/sql"
	"time"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/sqlutil"
)

type BookDTO struct {
	ID        int64
	AuthorID  int64
	BookName  string
	Title     sql.NullString
	Price     float64
	CreatedAt time.Time
}

func (dto *BookDTO) TransformToDomain() *domain.Book {
	return &domain.Book{
		ID:        dto.ID,
		AuthorID:  dto.AuthorID,
		BookName:  dto.BookName,
		Title:     dto.Title.String,
		Price:     dto.Price,
		CreatedAt: dto.CreatedAt,
	}
}

func (dto *BookDTO) TransformFromDomain(book *domain.Book) {
	dto.ID = book.ID
	dto.AuthorID = book.AuthorID
	dto.BookName = book.BookName
	dto.Title = sqlutil.NewNull(book.Title).(sql.NullString)
	dto.Price = book.Price
	dto.CreatedAt = book.CreatedAt
}

type DetailBookDTO struct {
	ID         int64
	AuthorID   int64
	AuthorName string
	BookName   string
	Title      string
	Price      float64
	CreatedAt  time.Time
}

func (dto *DetailBookDTO) TransformToDomain() *domain.DetailBook {
	return &domain.DetailBook{
		ID:         dto.ID,
		AuthorID:   dto.AuthorID,
		AuthorName: dto.AuthorName,
		BookName:   dto.BookName,
		Title:      dto.Title,
		Price:      dto.Price,
		CreatedAt:  dto.CreatedAt,
	}
}
