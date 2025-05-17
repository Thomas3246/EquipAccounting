package service

import (
	domain "github.com/Thomas3246/EquipAccounting/internal/domain/workstation"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type WorkStationService struct {
	repo database.WorkStationRepo
}

func NewWorkStationService(repo database.WorkStationRepo) *WorkStationService {
	return &WorkStationService{repo: repo}
}

func (s *WorkStationService) RegisterWorkStation(ws *domain.WorkStation) error {
	return s.repo.Create(ws)
}
