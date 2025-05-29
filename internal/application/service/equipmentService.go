package service

import (
	"context"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type EquipmentService struct {
	repo database.EquipmentRepo
}

func NewEquipmentService(repo database.EquipmentRepo) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (s *EquipmentService) GetAvailableEquipment(login string) ([]domain.EquipmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	equipment, err := s.repo.GetActiveEquipmentForUserLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}
