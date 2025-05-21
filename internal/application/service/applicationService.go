package service

import (
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

type ApplicationService struct {
	WorkStationService *WorkStationService
	UserService        *UserService
	RequestService     *RequestService
}

func NewAppService(db *sql.DB) *ApplicationService {
	workStationRepo := sqlite.NewWorkStationRepo(db)
	workStationService := NewWorkStationService(workStationRepo)

	userRepo := sqlite.NewUserRepo(db)
	userService := NewUserService(userRepo)

	requestRepo := sqlite.NewRequestRepo(db)
	requestService := NewRequestService(requestRepo)

	return &ApplicationService{
		WorkStationService: workStationService,
		UserService:        userService,
		RequestService:     requestService,
	}
}
