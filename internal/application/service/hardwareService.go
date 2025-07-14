package service

import (
	"context"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type HardwareService struct {
	repo database.HardwareRepo
}

func NewHardwareService(repo database.HardwareRepo) *HardwareService {
	return &HardwareService{repo: repo}
}

func (s *HardwareService) GetUnitsByType(hType string) ([]domain.Unit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if hType != "cpu" && hType != "gpu" && hType != "motherboard" {
		return nil, ErrInvalidParameter
	}

	units, err := s.repo.GetUnitsByType(ctx, hType)
	if err != nil {
		return nil, err
	}

	return units, nil
}

func (s *HardwareService) GetUnit(hType string, id int) (domain.Unit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if hType != "cpu" && hType != "gpu" && hType != "motherboard" {
		return nil, ErrInvalidParameter
	}

	unit, err := s.repo.GetUnit(ctx, hType, id)
	if err != nil {
		return nil, err
	}
	return unit, nil
}

func (s *HardwareService) UpdateName(hType string, id int, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if hType != "cpu" && hType != "gpu" && hType != "motherboard" {
		return ErrInvalidParameter
	}

	err := s.repo.UpdateName(ctx, hType, id, name)
	return err
}

func (s *HardwareService) NewUnit(hType string, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.NewUnit(ctx, hType, name)
	return err
}

func (s *HardwareService) DeleteUnit(hType string, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if hType != "cpu" && hType != "gpu" && hType != "motherboard" {
		return ErrInvalidParameter
	}

	err := s.repo.DeleteUnit(ctx, hType, id)
	return err
}
