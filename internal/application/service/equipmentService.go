package service

import (
	"context"
	"strconv"
	"strings"
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

func (s *EquipmentService) GetAvailableEquipment(cookieValue string) ([]domain.EquipmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return nil, ErrInvalidCookieParameter
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	if isAdmin == 1 {
		equipment, err := s.repo.GetActiveEquipment(ctx)
		if err != nil {
			return nil, err
		}
		return equipment, nil
	}

	if isAdmin == 0 {
		equipment, err := s.repo.GetActiveEquipmentForUserLogin(ctx, parts[0])
		if err != nil {
			return nil, err
		}
		return equipment, nil
	}

	return nil, ErrInvalidIsAdminValue
}
