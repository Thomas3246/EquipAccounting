package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type RequestRepo interface {
	GetAllActive(context.Context) ([]domain.Request, error)
	GetAllActiveDetail(context.Context) ([]domain.RequestView, error)
	GetAllActiveForUserDetail(context.Context, string) ([]domain.RequestView, error)
}
