package domain

import "time"

type DetailBook struct {
	ID         int       `gorm:"column:id" json:"id"`
	AuthorID   int       `gorm:"column:author_id" json:"author_id"`
	AuthorName string    `gorm:"column:author_name" json:"author_name"`
	BookName   string    `gorm:"column:book_name" json:"book_name"`
	Title      string    `gorm:"column:title" json:"title"`
	Price      int       `gorm:"column:price" json:"price"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}

type UpdateBookRequest struct {
	ID       int
	AuthorID int    `json:"author_id"`
	BookName string `json:"book_name"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
}

type CreateBookRequest struct {
	AuthorID  int       `json:"author_id"`
	BookName  string    `json:"book_name"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type Book struct {
	ID        int       `gorm:"column:id" json:"id"`
	AuthorID  int       `gorm:"column:author_id" json:"author_id"`
	BookName  string    `gorm:"column:book_name" json:"book_name"`
	Title     string    `gorm:"column:title" json:"title"`
	Price     int       `gorm:"column:price" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Book) TableName() string {
	return "books"
}
