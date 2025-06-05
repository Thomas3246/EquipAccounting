package service

import (
	"context"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type EquipmentDirectoryService struct {
	repo database.EquipmentDirectoryRepo
}

func NewEquipmentDirectoryService(repo database.EquipmentDirectoryRepo) *EquipmentDirectoryService {
	return &EquipmentDirectoryService{repo: repo}
}

func (s *EquipmentDirectoryService) GetEquipmentDirectoryTypes() ([]domain.EquipmentDirectoryType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	types, err := s.repo.GetEquipmentDirectoryTypes(ctx)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (s *EquipmentDirectoryService) GetEquipmentDirectoriesViewByFilter(eType int) ([]domain.EquipmentDirectoryView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	directories, err := s.repo.GetEquipmentDirectoriesViewByFilter(ctx, eType)
	if err != nil {
		return nil, err
	}

	return directories, nil
}
