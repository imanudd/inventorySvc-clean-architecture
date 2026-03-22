package dto

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
)

type CreateAuthorRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func (r *CreateAuthorRequest) TransformIn(c *gin.Context) (*domain.CreateAuthorRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	return &domain.CreateAuthorRequest{
		Name:        r.Name,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
	}, nil
}

type AddAuthorBookRequest struct {
	BookName string  `json:"book_name" validate:"required"`
	Title    string  `json:"title" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

func (r *AddAuthorBookRequest) TransformIn(c *gin.Context) (*domain.AddAuthorBookRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	return &domain.AddAuthorBookRequest{
		AuthorID: int64(authorID),
		BookName: r.BookName,
		Title:    r.Title,
		Price:    r.Price,
	}, nil
}

type CreateAuthorAndBookRequest struct {
	Author CreateAuthorRequest `json:"author" validate:"required"`
	Book   CreateBookRequest   `json:"book" validate:"required"`
}

func (r *CreateAuthorAndBookRequest) TransformIn(c *gin.Context) (*domain.CreateAuthorAndBookRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	authorDomain := &domain.CreateAuthorRequest{
		Name:        r.Author.Name,
		Email:       r.Author.Email,
		PhoneNumber: r.Author.PhoneNumber,
	}

	bookDomain := &domain.CreateBookRequest{
		AuthorID:  r.Book.AuthorID,
		BookName:  r.Book.BookName,
		Title:     r.Book.Title,
		Price:     r.Book.Price,
		CreatedAt: r.Book.CreatedAt,
	}

	return &domain.CreateAuthorAndBookRequest{
		Author: *authorDomain,
		Book:   *bookDomain,
	}, nil
}

type AuthorResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (r *AuthorResponse) TransformOut(author *domain.Author) {
	r.ID = author.ID
	r.Name = author.Name
	r.Email = author.Email
	r.PhoneNumber = author.PhoneNumber
}
