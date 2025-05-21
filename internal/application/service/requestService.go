package service

import (
	"context"
	"time"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/request"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type RequestService struct {
	repo database.RequestRepo
}

func NewRequestService(repo database.RequestRepo) *RequestService {
	return &RequestService{repo: repo}
}

func (s *RequestService) GetAllActive() ([]domain.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	requests, err := s.repo.GetAllActive(ctx)
	if err != nil {
		return nil, err
	}
	return requests, nil
}
