package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type DepartmentRepo interface {
	GetDepartmentsView(context.Context) ([]domain.DepartmentView, error)
}
