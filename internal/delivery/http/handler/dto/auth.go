package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) TransformIn(c *gin.Context) (*domain.LoginRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	return &domain.LoginRequest{
		Username: r.Username,
		Password: r.Password,
	}, nil
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (r *RegisterRequest) TransformIn(c *gin.Context) (*domain.RegisterRequest, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(r); err != nil {
		return nil, err
	}

	return &domain.RegisterRequest{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}, nil
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (r *LoginResponse) TransformOut(resp *domain.LoginResponse) {
	r.Username = resp.Username
	r.Token = resp.Token
}
