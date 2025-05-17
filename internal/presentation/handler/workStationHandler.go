package handler

import "github.com/Thomas3246/EquipAccounting/internal/application/service"

type WorkSationHandler struct {
	service service.WorkStationService
}

func NewWorkStationHandler(service *service.WorkStationService) *WorkSationHandler {
	return &WorkSationHandler{service: *service}
}
