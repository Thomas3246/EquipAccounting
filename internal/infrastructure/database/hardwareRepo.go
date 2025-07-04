package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type HardwareRepo interface {
	GetEquipmentHardware(context.Context, int) (domain.Hardware, error)
}
