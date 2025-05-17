package handler

import "github.com/Thomas3246/EquipAccounting/internal/application/service"

type AppHandler struct {
	WorkSationHandler *WorkSationHandler
}

func NewAppHandler(service *service.AppclicationService) *AppHandler {
	return &AppHandler{
		WorkSationHandler: NewWorkStationHandler(service.WorkStationService),
	}
}
