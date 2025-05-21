package handler

import (
	"github.com/Thomas3246/EquipAccounting/internal/application/service"
)

type AppHandler struct {
	WorkStationHandler *WorkStationHandler
	UserHandler        *UserHandler
	RequestHandler     *RequestHandler
}

func NewAppHandler(service *service.ApplicationService) *AppHandler {
	return &AppHandler{
		WorkStationHandler: NewWorkStationHandler(service.WorkStationService),
		UserHandler:        NewUserHandler(service.UserService),
		RequestHandler:     NewRequestHandler(service.RequestService),
	}
}
