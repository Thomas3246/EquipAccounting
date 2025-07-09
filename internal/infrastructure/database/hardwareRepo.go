package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type HardwareRepo interface {
	GetEquipmentHardware(context.Context, int) (domain.Hardware, error)
	GetUnitsByType(context.Context, string) ([]domain.Unit, error)
	GetUnit(context.Context, string, int) (domain.Unit, error)
	UpdateName(context.Context, string, int, string) error
}
