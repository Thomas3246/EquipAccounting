package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type UserRepo interface {
	Create(context.Context, *domain.User) error
	GetByLogin(context.Context, string) (*domain.User, error)
}
