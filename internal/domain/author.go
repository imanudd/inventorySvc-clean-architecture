package domain

type CreateAuthorAndBookRequest struct {
	Author CreateAuthorRequest
	Book   CreateBookRequest
}

type AddAuthorBookRequest struct {
	AuthorID int64
	BookName string
	Title    string
	Price    float64
}

type CreateAuthorRequest struct {
	Name        string
	Email       string
	PhoneNumber string
}

type Author struct {
	ID          int64
	Name        string
	Email       string
	PhoneNumber string
}

func (Author) TableName() string {
	return "authors"
}
