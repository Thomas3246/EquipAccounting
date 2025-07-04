package service

import "github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"

type HardwareService struct {
	repo database.HardwareRepo
}

func NewHardwareService(repo database.HardwareRepo) *HardwareService {
	return &HardwareService{repo: repo}
}
