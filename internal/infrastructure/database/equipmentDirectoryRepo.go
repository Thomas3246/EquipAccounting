package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentDirectoryRepo interface {
	GetEquipmentDirectoryTypes(context.Context) ([]domain.EquipmentDirectoryType, error)
	GetEquipmentDirectoriesViewByFilter(context.Context, int) ([]domain.EquipmentDirectoryView, error)
}
