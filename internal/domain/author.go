package domain

type AddAuthorBookRequest struct {
	AuthorID int    `json:"-"`
	BookName string `json:"book_name"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
}

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
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
