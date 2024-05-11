package domain

type CreateAuthorAndBookRequest struct {
	Author CreateAuthorRequest `json:"author" validate:"required"`
	Book   CreateBookRequest   `json:"book" validate:"required"`
}

type AddAuthorBookRequest struct {
	AuthorID int    `json:"-"`
	BookName string `json:"book_name" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Price    int    `json:"price" validate:"required"`
}

type CreateAuthorRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Author struct {
	ID          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	PhoneNumber string `gorm:"column:phone_number"`
}

func (Author) TableName() string {
	return "authors"
}
