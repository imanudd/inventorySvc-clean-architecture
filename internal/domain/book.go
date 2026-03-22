package domain

import "time"

type CreateDetailBook struct {
	Id       int64
	BookName string
	Title    string
	Price    float64
	Author   Author
}

type DetailBook struct {
	ID         int64
	AuthorID   int64
	AuthorName string
	BookName   string
	Title      string
	Price      float64
	CreatedAt  time.Time
}

type UpdateBookRequest struct {
	ID       int64
	AuthorID int64
	BookName string
	Title    string
	Price    float64
}

type CreateBookRequest struct {
	AuthorID  int64
	BookName  string
	Title     string
	Price     float64
	CreatedAt time.Time
}

type Book struct {
	ID        int64
	AuthorID  int64
	BookName  string
	Title     string
	Price     float64
	CreatedAt time.Time
}

func (Book) TableName() string {
	return "books"
}
