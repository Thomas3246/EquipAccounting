package service

import (
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

type ApplicationService struct {
	UserService    *UserService
	RequestService *RequestService
}

func NewAppService(db *sql.DB) *ApplicationService {

	userRepo := sqlite.NewUserRepo(db)
	userService := NewUserService(userRepo)

	requestRepo := sqlite.NewRequestRepo(db)
	requestService := NewRequestService(requestRepo)

	return &ApplicationService{
		UserService:    userService,
		RequestService: requestService,
	}
}
