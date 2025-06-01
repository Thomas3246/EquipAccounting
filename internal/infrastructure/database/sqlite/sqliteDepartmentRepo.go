package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type DepartmentRepo struct {
	db *sql.DB
}

func NewDepartmentRepo(db *sql.DB) *DepartmentRepo {
	return &DepartmentRepo{db: db}
}

func (r *DepartmentRepo) GetDepartmentsView(ctx context.Context) (departments []domain.DepartmentView, err error) {
	query := `SELECT dep.Id, dep.name, div.name
			  FROM department as dep
			  INNER JOIN departmentDivisions as div ON dep.division = div.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d := domain.DepartmentView{}
		err = rows.Scan(&d.Id, &d.Name, &d.DivisionName)
		if err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}
	return departments, nil
}
