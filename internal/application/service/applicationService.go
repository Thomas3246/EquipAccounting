package service

import (
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

type AppclicationService struct {
	WorkStationService *WorkStationService
}

func NewAppService(db *sql.DB) *AppclicationService {
	workStationRepo := sqlite.NewWorkStationRepo(db)
	workStationService := NewWorkStationService(workStationRepo)

	return &AppclicationService{
		WorkStationService: workStationService,
	}
}
