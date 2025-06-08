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
	GetEquipmentById(context.Context, int) (domain.Equipment, error)
	GetEquipmentByInvNum(context.Context, string) (domain.Equipment, error)
	UpdateEquipment(context.Context, domain.Equipment) error
	DeleteEquipment(context.Context, int) error
}
