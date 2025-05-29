package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentRepo struct {
	db *sql.DB
}

func NewEquipmentRepo(db *sql.DB) *EquipmentRepo {
	return &EquipmentRepo{db: db}
}

func (r *EquipmentRepo) GetActiveEquipmentForUserLogin(ctx context.Context, login string) (equipmentList []domain.EquipmentView, err error) {
	quey := `SELECT e.id, e.invNum, ed.name
			 FROM equipment as e
			 INNER JOIN equipDirectory as ed ON e.directory = ed.id
			 INNER JOIN users as u ON u.department = e.department
			 WHERE u.login = ? AND e.status = 1`

	rows, err := r.db.QueryContext(ctx, quey, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		eq := domain.EquipmentView{}
		err = rows.Scan(&eq.Id, &eq.InvNum, &eq.Directory)
		if err != nil {
			return nil, err
		}
		equipmentList = append(equipmentList, eq)
	}
	return equipmentList, nil
}
