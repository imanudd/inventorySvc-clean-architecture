package dto

import (
	"database/sql"

	"github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	"github.com/imanudd/inventorySvc-clean-architecture/pkg/sqlutil"
)

type AuthorDTO struct {
	ID          int64
	Name        string
	Email       string
	PhoneNumber sql.NullString
}

func (dto *AuthorDTO) TransformToDomain() *domain.Author {
	return &domain.Author{
		ID:          dto.ID,
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber.String,
	}
}

func (dto *AuthorDTO) TransformFromDomain(author *domain.Author) {
	dto.ID = author.ID
	dto.Name = author.Name
	dto.Email = author.Email
	dto.PhoneNumber = sqlutil.NewNull(author.PhoneNumber).(sql.NullString)
}
