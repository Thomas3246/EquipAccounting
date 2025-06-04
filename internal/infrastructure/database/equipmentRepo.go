package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentRepo interface {
	GetActiveEquipmentForUserLogin(context.Context, string) ([]domain.EquipmentView, error)
	GetActiveEquipment(context.Context) ([]domain.EquipmentView, error)
	GetEquipmentStates(context.Context) ([]domain.EquipmentState, error)
	GetEquipmentViewByFilter(context.Context, int, int) ([]domain.EquipmentView, error)
}
