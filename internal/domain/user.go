package domain

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetByUsernameOrEmail struct {
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type User struct {
	ID       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
