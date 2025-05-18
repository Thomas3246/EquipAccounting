package handler

import "github.com/Thomas3246/EquipAccounting/internal/application/service"

type WorkStationHandler struct {
	service service.WorkStationService
}

func NewWorkStationHandler(service *service.WorkStationService) *WorkStationHandler {
	return &WorkStationHandler{service: *service}
}
