package sqlite

import (
	"database/sql"
	"fmt"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/workstation"
)

type WorkStationRepo struct {
	db *sql.DB
}

func NewWorkStationRepo(db *sql.DB) *WorkStationRepo {
	return &WorkStationRepo{db: db}
}

func (r *WorkStationRepo) Create(eq *domain.WorkStation) error {
	fmt.Println("Added")
	return nil
}
