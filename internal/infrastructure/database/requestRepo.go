package database

import (
	"context"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/request"
)

type RequestRepo interface {
	GetAllActive(context.Context) ([]domain.Request, error)
}
