package service

import (
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

type ApplicationService struct {
	WorkStationService *WorkStationService
	UserService        *UserService
}

func NewAppService(db *sql.DB) *ApplicationService {
	workStationRepo := sqlite.NewWorkStationRepo(db)
	workStationService := NewWorkStationService(workStationRepo)

	userRepo := sqlite.NewUserRepo(db)
	UserService := NewUserService(userRepo)

	return &ApplicationService{
		WorkStationService: workStationService,
		UserService:        UserService,
	}
}
