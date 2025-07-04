package handler

import "github.com/Thomas3246/EquipAccounting/internal/application/service"

type HardwareHandler struct {
	hardwareService service.HardwareService
}

func NewHardwareHandler(hardwareService service.HardwareService) *HardwareHandler {
	return &HardwareHandler{hardwareService: hardwareService}
}
