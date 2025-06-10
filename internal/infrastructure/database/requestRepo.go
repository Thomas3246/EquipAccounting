package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type RequestRepo interface {
	GetAllActive(context.Context) ([]domain.Request, error)
	GetAllActiveDetail(context.Context) ([]domain.RequestView, error)
	GetAllClosedDetail(context.Context) ([]domain.RequestView, error)
	GetAllActiveForUserDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllClosedForUserDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllUserActiveDetail(context.Context, string) ([]domain.RequestView, error)
	GetAllUserClosedDetail(context.Context, string) ([]domain.RequestView, error)
	GetRequestTypes(context.Context) ([]domain.RequestType, error)
	AddRequest(context.Context, domain.Request) error
	GetRequestById(context.Context, int) (domain.Request, error)
	UpdateDescription(context.Context, int, string) error
	UpdateRequest(context.Context, domain.Request) error
	GetResults(context.Context) ([]domain.RequestResult, error)
	GetRequestsWithEquipment(context.Context, int) ([]domain.Request, error)
	CloseRequest(context.Context, int, int, string, string) error
	GetReportData(context.Context, int, string) (domain.RequestReport, error)
}
