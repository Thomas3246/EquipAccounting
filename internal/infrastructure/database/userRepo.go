package database

import (
	"context"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/user"
)

type UserRepo interface {
	Create(*domain.User) error
	GetByLogin(context.Context, string) (*domain.User, error)
}
