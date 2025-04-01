package usecase

import (
	"github.com/imanudd/inventorySvc-clean-architecture/config"
	"github.com/imanudd/inventorySvc-clean-architecture/internal/repository"
)

type Usecase struct {
	AuthUseCase   AuthUseCaseImpl
	BookUseCase   BookUseCaseImpl
	AuthorUseCase AuthorUseCaseImpl
}

func NewUsecase(cfg *config.MainConfig, repository repository.RepositoryImpl) Usecase {
	return Usecase{
		AuthUseCase:   NewAuthUseCase(cfg, repository),
		BookUseCase:   NewBookUseCase(cfg, repository),
		AuthorUseCase: NewAuthorUseCase(cfg, repository),
	}
}

func (u *Usecase) GetAuthUseCase() AuthUseCaseImpl {
	return u.AuthUseCase
}

func (u *Usecase) GetBookUseCase() BookUseCaseImpl {
	return u.BookUseCase
}

func (u *Usecase) GetAuthorUseCase() AuthorUseCaseImpl {
	return u.AuthorUseCase
}
