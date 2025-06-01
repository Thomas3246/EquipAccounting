package service

import (
	"context"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type DepartmentService struct {
	repo database.DepartmentRepo
}

func NewDepartmentService(repo database.DepartmentRepo) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) GetDepartmentsView() ([]domain.DepartmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	departments, err := s.repo.GetDepartmentsView(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}
