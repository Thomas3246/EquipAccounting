package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentDirectoryRepo interface {
	GetEquipmentDirectoryTypes(context.Context) ([]domain.EquipmentDirectoryType, error)
	GetEquipmentDirectoriesViewByFilter(context.Context, int) ([]domain.EquipmentDirectoryView, error)
	GetDirectory(context.Context, int) (domain.EquipmentDirectory, error)
	UpdateDirectory(context.Context, domain.EquipmentDirectory) error
	NewDirectory(context.Context, domain.EquipmentDirectory) error
	DeleteDirectory(context.Context, int) error
}
