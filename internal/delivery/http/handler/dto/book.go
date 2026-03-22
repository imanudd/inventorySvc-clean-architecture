package dto

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/delivery/http/helper"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
)

type CreateBookRequest struct {
	AuthorID  int64     `json:"author_id" validate:"required"`
	BookName  string    `json:"book_name" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *CreateBookRequest) TransformIn(c *gin.Context) (*domain.CreateBookRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	return &domain.CreateBookRequest{
		AuthorID:  r.AuthorID,
		BookName:  r.BookName,
		Title:     r.Title,
		Price:     r.Price,
		CreatedAt: r.CreatedAt,
	}, nil
}

type UpdateBookRequest struct {
	ID       int64   `json:"-"`
	AuthorID int64   `json:"author_id" validate:"required"`
	BookName string  `json:"book_name" validate:"required"`
	Title    string  `json:"title" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

func (r *UpdateBookRequest) TransformIn(c *gin.Context) (*domain.UpdateBookRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "error bad request")
		return nil, err
	}

	return &domain.UpdateBookRequest{
		ID:       int64(id),
		AuthorID: r.AuthorID,
		BookName: r.BookName,
		Title:    r.Title,
		Price:    r.Price,
	}, nil
}

type BookResponse struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	BookName  string    `json:"book_name"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *BookResponse) TransformOut(book *domain.Book) {
	r.ID = book.ID
	r.AuthorID = book.AuthorID
	r.BookName = book.BookName
	r.Title = book.Title
	r.Price = book.Price
	r.CreatedAt = book.CreatedAt
}

type DetailBookResponse struct {
	ID         int64     `json:"id"`
	AuthorID   int64     `json:"author_id"`
	AuthorName string    `json:"author_name"`
	BookName   string    `json:"book_name"`
	Title      string    `json:"title"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
}

func (r *DetailBookResponse) TransformOut(detail *domain.DetailBook) {
	r.ID = detail.ID
	r.AuthorID = detail.AuthorID
	r.AuthorName = detail.AuthorName
	r.BookName = detail.BookName
	r.Title = detail.Title
	r.Price = detail.Price
	r.CreatedAt = detail.CreatedAt
}

type CreateDetailBook struct {
	Id       int64          `json:"id"`
	BookName string         `json:"book_name"`
	Title    string         `json:"title"`
	Price    float64        `json:"price"`
	Author   AuthorResponse `json:"author"`
}

func (r *CreateDetailBook) TransformOut(detail *domain.CreateDetailBook) {
	r.Id = detail.Id
	r.BookName = detail.BookName
	r.Title = detail.Title
	r.Price = detail.Price

	r.Author = AuthorResponse{}
	r.Author.TransformOut(&detail.Author)
}
