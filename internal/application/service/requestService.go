package service

import (
	"context"
	"database/sql"
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

func (s *RequestService) GetAllUserActive(cookieValue string, requestedLogin string) (requests []domain.RequestView, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return nil, err
	}

	if requestedLogin == parts[0] || parts[1] == "1" {
		requests, err = s.repo.GetAllUserActiveDetail(ctx, requestedLogin)
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	return nil, ErrNoAccess

}

func (s *RequestService) GetAllClosed(cookieValue string) (requests []domain.RequestView, err error) {
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
		requests, err = s.repo.GetAllClosedDetail(ctx)
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	if isAdmin == 0 {
		requests, err = s.repo.GetAllClosedForUserDetail(ctx, parts[0])
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	return nil, ErrInvalidCookieParameter
}

func (s *RequestService) GetAllUserClosed(cookieValue string, requestedLogin string) (requests []domain.RequestView, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return nil, err
	}

	if requestedLogin == parts[0] || parts[1] == "1" {
		requests, err = s.repo.GetAllUserClosedDetail(ctx, requestedLogin)
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	return nil, ErrNoAccess

}

func (s *RequestService) GetRequestTypes() (types []domain.RequestType, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	types, err = s.repo.GetRequestTypes(ctx)
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (s *RequestService) NewRequest(rType int, descr string, author int, equipment int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := domain.Request{
		Type:        rType,
		Description: descr,
		Author:      author,
		Status:      1,
		CreatedAt:   time.Now().Format("2006-01-02 15:04"),
		Equipment:   equipment,
	}

	err := s.repo.AddRequest(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (s *RequestService) GetRequestById(id int) (domain.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request, err := s.repo.GetRequestById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return request, ErrNotFound
		}
		return request, err
	}

	return request, nil
}

func (s *RequestService) EditDescription(reqId int, descr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.repo.UpdateDescription(ctx, reqId, descr)
}

func (s *RequestService) EditRequest(request domain.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.repo.UpdateRequest(ctx, request)
}
