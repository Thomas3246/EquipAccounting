package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type RequestService struct {
	repo database.RequestRepo
}

func NewRequestService(repo database.RequestRepo) *RequestService {
	return &RequestService{repo: repo}
}

func (s *RequestService) GetAllActive(cookieValue string) (requests []domain.RequestView, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return nil, err
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	if isAdmin == 1 {
		requests, err = s.repo.GetAllActiveDetail(ctx)
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	if isAdmin == 0 {
		requests, err = s.repo.GetAllActiveForUserDetail(ctx, parts[0])
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	return nil, ErrInvalidCookieParameter
}
