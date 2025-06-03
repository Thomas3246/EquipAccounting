package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type UserRepo interface {
	Create(context.Context, *domain.User) error
	GetByLogin(context.Context, string) (*domain.User, error)
	GetUsersView(context.Context) ([]domain.ViewUser, error)
	Delete(context.Context, int) error
	GetById(context.Context, int) (domain.User, error)
	New(context.Context, *domain.User) error
	ChangeUserPassword(context.Context, int, string) error
	UpdateUserData(context.Context, domain.User) error
}
