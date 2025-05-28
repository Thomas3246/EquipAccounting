package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type RequestRepo interface {
	GetAllActive(context.Context) ([]domain.Request, error)
	GetAllActiveDetail(context.Context) ([]domain.RequestView, error)
	GetAllClosedDetail(context.Context) ([]domain.RequestView, error)
	GetAllActiveForUserDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllClosedForUserDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllUserActiveDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllUserClosedDetail(context.Context, string) ([]domain.RequestView, error)
}
