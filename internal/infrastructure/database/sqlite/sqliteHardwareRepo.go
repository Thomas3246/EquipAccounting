package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type HardwareRepo struct {
	db *sql.DB
}

func NewHardwareRepo(db *sql.DB) *HardwareRepo {
	return &HardwareRepo{db: db}
}

func (r *HardwareRepo) GetEquipmentHardware(ctx context.Context, id int) (hardware domain.Hardware, err error) {
	query := `SELECT h.id, h.RAM, h.storage, cpu.id, cpu.name, COALESCE(gpu.id, 0) AS gpuid, COALESCE(gpu.name, '') AS gpuname, mb.id, mb.name
			  FROM hardware AS h
			  INNER JOIN cpu ON h.cpuid = cpu.id
			  LEFT JOIN gpu ON h.gpuid = gpu.id
			  INNER JOIN motherboard AS mb ON h.motherboardid = mb.id
			  WHERE h.equipDirectoryId = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	err = row.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Hardware{}, nil
		}
		return domain.Hardware{}, err
	}

	hardware.EquipmentDirectoryId = id
	err = row.Scan(&hardware.Id, &hardware.RAM, &hardware.CPU.Id, &hardware.CPU.Name, &hardware.GPU.Id, &hardware.GPU.Name, &hardware.Motherboard.Id, &hardware.Motherboard.Name)
	if err != nil {
		return domain.Hardware{}, err
	}
	return hardware, nil
}

func (r *HardwareRepo) GetUnitsByType(ctx context.Context, hType string) (units []domain.Unit, err error) {
	query := fmt.Sprintf("SELECT id, name FROM %s", hType)
	rows, err := r.db.QueryContext(ctx, query, hType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		switch hType {
		case "cpu":
			units = append(units, domain.CPU{Id: id, Name: name})
		case "gpu":
			units = append(units, domain.GPU{Id: id, Name: name})
		case "motherboard":
			units = append(units, domain.Motherboard{Id: id, Name: name})
		}
	}

	return units, nil
}

func (r *HardwareRepo) GetUnit(ctx context.Context, hType string, id int) (domain.Unit, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = ?", hType)

	row := r.db.QueryRowContext(ctx, query, id)

	switch hType {
	case "cpu":
		var cpu domain.CPU
		err := row.Scan(&cpu.Id, &cpu.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return cpu, nil

	case "gpu":
		var gpu domain.GPU
		err := row.Scan(&gpu.Id, &gpu.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return gpu, nil

	case "motherboard":
		var mb domain.Motherboard
		err := row.Scan(&mb.Id, &mb.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return mb, nil

	default:
		return nil, errors.New("unsupported unit type")
	}
}

func (r *HardwareRepo) UpdateName(ctx context.Context, hType string, id int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = ? WHERE id = ?", hType)

	_, err := r.db.ExecContext(ctx, query, name, id)
	return err
}
