package service

import (
	"context"
	"database/sql"
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

func (s *EquipmentDirectoryService) GetEquipmentDirectory(id int) (domain.EquipmentDirectory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	directory, err := s.repo.GetDirectory(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.EquipmentDirectory{}, ErrRequestNotFound
		}
		return domain.EquipmentDirectory{}, err
	}
	return directory, nil
}

func (s *EquipmentDirectoryService) UpdateDirectory(directory domain.EquipmentDirectory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.UpdateDirectory(ctx, directory)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentDirectoryService) NewDirectory(directory domain.EquipmentDirectory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.NewDirectory(ctx, directory)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentDirectoryService) DeleteDirectory(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.DeleteDirectory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
