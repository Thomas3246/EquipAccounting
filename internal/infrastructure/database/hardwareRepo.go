package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type HardwareRepo interface {
	GetUnitsByType(context.Context, string) ([]domain.Unit, error)
	GetUnit(context.Context, string, int) (domain.Unit, error)
	UpdateName(context.Context, string, int, string) error
	NewUnit(context.Context, string, string) error
	DeleteUnit(context.Context, string, int) error
}
