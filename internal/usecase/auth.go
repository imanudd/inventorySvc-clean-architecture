package usecase

import (
	"context"
	"errors"

	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/auth"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseImpl interface {
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error)
	Register(ctx context.Context, req *domain.RegisterRequest) (err error)
}

type authUseCase struct {
	cfg  *config.MainConfig
	repo repository.RepositoryImpl
}

func NewAuthUseCase(cfg *config.MainConfig, repo repository.RepositoryImpl) AuthUseCaseImpl {
	return &authUseCase{
		cfg:  cfg,
		repo: repo,
	}
}

func (a *authUseCase) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	if err := validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	user, err := a.repo.GetUserRepo().GetByUsernameOrEmail(ctx, &domain.GetByUsernameOrEmail{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user is not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	auth := auth.NewAuth(a.cfg)
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Username: req.Username,
		Token:    token,
	}, nil
}

func (a *authUseCase) Register(ctx context.Context, req *domain.RegisterRequest) (err error) {
	if err := validator.ValidateStruct(req); err != nil {
		return err
	}

	user, err := a.repo.GetUserRepo().GetByUsernameOrEmail(ctx, &domain.GetByUsernameOrEmail{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return
	}

	if user != nil {
		return errors.New("user is already exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error when hashing password")
	}

	return a.repo.GetUserRepo().RegisterUser(ctx, &domain.User{
		Username: req.Username,
		Password: string(hash),
		Email:    req.Email,
	})

}
