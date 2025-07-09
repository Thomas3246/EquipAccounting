package handler

import (
	"github.com/Thomas3246/EquipAccounting/internal/application/service"
)

type AppHandler struct {
	UserHandler               *UserHandler
	RequestHandler            *RequestHandler
	EquipmentHandler          *EquipmentHandler
	EquipmentDirectoryHandler *EquipmentDirectoryHandler
	DocumentHandler           *DocumentHandler
	HardwareHandler           *HardwareHandler
}

func NewAppHandler(service *service.ApplicationService) *AppHandler {
	return &AppHandler{
		UserHandler:               NewUserHandler(service.UserService, service.DepartmentService),
		RequestHandler:            NewRequestHandler(service.RequestService, service.UserService, service.EquipmentService, service.DocumentService),
		EquipmentHandler:          NewEquipmentHandler(service.EquipmentService, service.EquipmentDirectoryService, service.DepartmentService),
		EquipmentDirectoryHandler: NewEquipmentDirectoryHandler(service.EquipmentDirectoryService, service.EquipmentService),
		DocumentHandler:           NewDocumentHandler(service.DocumentService, service.UserService),
		HardwareHandler:           NewHardwareHandler(service.HardwareService),
	}
}
