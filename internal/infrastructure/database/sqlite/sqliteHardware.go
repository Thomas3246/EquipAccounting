package sqlite

import (
	"context"
	"database/sql"

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
