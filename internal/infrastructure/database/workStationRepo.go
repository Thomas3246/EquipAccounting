package database

import domain "github.com/Thomas3246/EquipAccounting/internal/domain/workstation"

type WorkStationRepo interface {
	Create(*domain.WorkStation) error
}
