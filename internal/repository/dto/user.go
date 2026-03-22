package dto

import (
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
)

type UserDTO struct {
	ID       int64
	Username string
	Email    string
	Password string
}

func (dto *UserDTO) TransformToDomain() *domain.User {
	return &domain.User{
		ID:       int(dto.ID),
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (dto *UserDTO) TransformFromDomain(user *domain.User) {
	dto.ID = int64(user.ID)
	dto.Username = user.Username
	dto.Email = user.Email
	dto.Password = user.Password
}
