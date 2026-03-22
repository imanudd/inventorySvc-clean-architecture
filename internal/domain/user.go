package domain

type LoginResponse struct {
	Username string
	Token    string
}

type LoginRequest struct {
	Username string
	Password string
}

type GetByUsernameOrEmail struct {
	Username string
	Email    string
}

type RegisterRequest struct {
	Username string
	Password string
	Email    string
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (User) TableName() string {
	return "users"
}
