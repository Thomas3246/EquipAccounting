package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentRepo struct {
	db *sql.DB
}

func NewEquipmentRepo(db *sql.DB) *EquipmentRepo {
	return &EquipmentRepo{db: db}
}

func (r *EquipmentRepo) GetActiveEquipment(ctx context.Context) (equipmentList []domain.EquipmentView, err error) {
	query := `SELECT e.id, e.invNum, ed.name
			 FROM equipment as e
			 INNER JOIN equipDirectory as ed ON e.directory = ed.id
			 WHERE e.status = 1`
	rows, err := r.db.QueryContext(ctx, query)
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

func (r *EquipmentRepo) GetEquipmentStates(ctx context.Context) (states []domain.EquipmentState, err error) {
	query := `SELECT id, name FROM equipStatus`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		state := domain.EquipmentState{}
		err = rows.Scan(&state.Id, &state.Name)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}

func (r *EquipmentRepo) GetEquipmentViewByFilter(ctx context.Context, department, state int) (equipment []domain.EquipmentView, err error) {
	query := `SELECT e.id, e.invNum, e.purchDate, e.regDate, 
			  	  COALESCE(e.decomDate, '') AS decomDate, dir.name || " " ||  COALESCE(dir.releaseYear, ''), dep.name || " - " || depdiv.name, es.Name
			  FROM equipment AS e
			  INNER JOIN equipDirectory AS dir ON e.directory = dir.id
			  INNER JOIN department AS dep ON e.department = dep.id
			  INNER JOIN departmentDivisions AS depdiv ON dep.division = depdiv.id
			  INNER JOIN equipStatus AS es ON e.status = es.id
			  WHERE 1=1`

	args := []any{}

	argPos := 1
	if department > 0 {
		query += fmt.Sprintf(" AND e.department = $%d", argPos)
		args = append(args, department)
		argPos++
	}

	if state > 0 {
		query += fmt.Sprintf(" AND e.status = $%d", argPos)
		args = append(args, state)
		argPos++
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := domain.EquipmentView{}
		err = rows.Scan(&e.Id, &e.InvNum, &e.PurchDate, &e.RegDate, &e.DecomDate, &e.Directory, &e.Department, &e.Status)
		if err != nil {
			return nil, err
		}
		equipment = append(equipment, e)
	}

	return equipment, nil
}

func (r *EquipmentRepo) GetEquipmentById(ctx context.Context, id int) (eq domain.Equipment, err error) {
	query := `SELECT id, invNum, purchDate, regDate, COALESCE(decomDate, '') AS  decomDate, directory, department, status, 
				COALESCE(cpu, 0) AS cpu, COALESCE(gpu, 0) AS gpu, COALESCE(motherboard, 0) AS motherboard, COALESCE(ram, 0) AS ram, COALESCE(storage, 0) AS storage  
			  FROM equipment
			  WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return domain.Equipment{}, err
	}

	err = row.Scan(&eq.Id, &eq.InvNum, &eq.PurchDate, &eq.RegDate, &eq.DecomDate, &eq.DirectoryId, &eq.DepartmentId, &eq.StatusId, &eq.CPU.Id, &eq.GPU.Id, &eq.Motherboard.Id, &eq.RAM, &eq.Storage)
	if err != nil {
		return domain.Equipment{}, err
	}
	return eq, nil
}

func (r *EquipmentRepo) GetEquipmentByInvNum(ctx context.Context, invNum string) (eq domain.Equipment, err error) {
	query := `SELECT id, invNum, purchDate, regDate, COALESCE(decomDate, '') AS  decomDate, directory, department, status
			  FROM equipment
			  WHERE invNum = ?`

	row := r.db.QueryRowContext(ctx, query, invNum)
	if row.Err() != nil {
		return domain.Equipment{}, err
	}

	err = row.Scan(&eq.Id, &eq.InvNum, &eq.PurchDate, &eq.RegDate, &eq.DecomDate, &eq.DirectoryId, &eq.DepartmentId, &eq.StatusId)
	if err != nil {
		return domain.Equipment{}, err
	}
	return eq, nil
}

func (r *EquipmentRepo) UpdateEquipment(ctx context.Context, equipment domain.Equipment) error {
	query := `UPDATE equipment
			  SET invNum = ?, directory = ?, department = ?, status = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, equipment.InvNum, equipment.DirectoryId, equipment.DepartmentId, equipment.StatusId, equipment.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) UpdatePC(ctx context.Context, equipment domain.Equipment) error {
	query := `UPDATE equipment 
			  SET invNum = ?, directory = ?, department = ?, status = ?, ram = ?, storage = ?, cpu = ?, gpu = ?, motherboard = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(
		ctx,
		query,
		equipment.InvNum,
		equipment.DirectoryId,
		equipment.DepartmentId,
		equipment.StatusId,
		equipment.RAM,
		equipment.Storage,
		equipment.CPU.Id,
		equipment.GPU.Id,
		equipment.Motherboard.Id,
		equipment.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) AddEquipment(ctx context.Context, equipment domain.Equipment) error {
	query := `INSERT INTO equipment (invNum, purchDate, regDate, directory, department, status)
			  VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, equipment.InvNum, equipment.PurchDate, equipment.RegDate, equipment.DirectoryId, equipment.DepartmentId, equipment.StatusId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) DeleteEquipment(ctx context.Context, id int) error {
	query := `DELETE FROM equipment WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) ChangeEquipStatus(ctx context.Context, equipId int, status int) error {
	query := `UPDATE equipment 
			  SET status = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, status, equipId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) DecomEquipment(ctx context.Context, equipId int, decomDate string) error {
	query := `UPDATE equipment
			  SET decomDate = ?, status = 2
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, decomDate, equipId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentRepo) GetAllEquipment(ctx context.Context) (equipment []domain.EquipmentView, err error) {
	query := `SELECT e.id, e.invNum, e.purchDate, e.regDate, 
			  	  COALESCE(e.decomDate, '') AS decomDate, dir.name || " " ||  COALESCE(dir.releaseYear, ''), dep.name || " - " || depdiv.name, es.Name
			  FROM equipment AS e
			  INNER JOIN equipDirectory AS dir ON e.directory = dir.id
			  INNER JOIN department AS dep ON e.department = dep.id
			  INNER JOIN departmentDivisions AS depdiv ON dep.division = depdiv.id
			  INNER JOIN equipStatus AS es ON e.status = es.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := domain.EquipmentView{}
		err = rows.Scan(&e.Id, &e.InvNum, &e.PurchDate, &e.RegDate, &e.DecomDate, &e.Directory, &e.Department, &e.Status)
		if err != nil {
			return nil, err
		}
		equipment = append(equipment, e)
	}
	return equipment, nil
}
