package handler

import (
	"github.com/Thomas3246/EquipAccounting/internal/application/service"
)

type AppHandler struct {
	UserHandler      *UserHandler
	RequestHandler   *RequestHandler
	EquipmantHandler *EquipmentHandler
}

func NewAppHandler(service *service.ApplicationService) *AppHandler {
	return &AppHandler{
		UserHandler:      NewUserHandler(service.UserService, service.DepartmentService),
		RequestHandler:   NewRequestHandler(service.RequestService, service.UserService, service.EquipmentService),
		EquipmantHandler: NewEquipmentHandler(service.EquipmentService, service.DepartmentService),
	}
}
