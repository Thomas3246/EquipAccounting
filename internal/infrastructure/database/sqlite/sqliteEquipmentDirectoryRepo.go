package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type EquipmentDirectoryRepo struct {
	db *sql.DB
}

func NewEquipmentDirectoryRepo(db *sql.DB) *EquipmentDirectoryRepo {
	return &EquipmentDirectoryRepo{db: db}
}

func (r *EquipmentDirectoryRepo) GetEquipmentDirectoryTypes(ctx context.Context) (types []domain.EquipmentDirectoryType, err error) {
	query := `SELECT * FROM equipType`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := domain.EquipmentDirectoryType{}
		err = rows.Scan(&t.Id, &t.Name)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, nil
}

func (r *EquipmentDirectoryRepo) GetEquipmentDirectoriesViewByFilter(ctx context.Context, eType int) (directories []domain.EquipmentDirectoryView, err error) {
	query := `SELECT ed.id, ed.name, COALESCE(ed.releaseYear, '') AS releaseYear, et.name
			  FROM equipDirectory AS ed
			  INNER JOIN equipType AS et ON ed.type = et.id
			  WHERE 1=1`

	args := []any{}

	argPos := 1
	if eType > 0 {
		query += fmt.Sprintf(" AND ed.type = $%d", argPos)
		args = append(args, eType)
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
		d := domain.EquipmentDirectoryView{}
		err = rows.Scan(&d.Id, &d.Name, &d.ReleaseYear, &d.Type)
		if err != nil {
			return nil, err
		}
		directories = append(directories, d)
	}

	return directories, nil
}

func (r *EquipmentDirectoryRepo) GetDirectory(ctx context.Context, id int) (directory domain.EquipmentDirectory, err error) {
	query := `SELECT id, name, COALESCE(releaseYear, '') AS releaseYear, type FROM equipDirectory WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return domain.EquipmentDirectory{}, row.Err()
	}

	err = row.Scan(&directory.Id, &directory.Name, &directory.ReleaseYear, &directory.TypeId)
	if err != nil {
		return domain.EquipmentDirectory{}, err
	}
	return directory, nil
}

func (r *EquipmentDirectoryRepo) UpdateDirectory(ctx context.Context, directory domain.EquipmentDirectory) error {
	query := `UPDATE equipDirectory
			  SET name = ?, releaseYear = ?, type = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, directory.Name, directory.ReleaseYear, directory.TypeId, directory.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentDirectoryRepo) NewDirectory(ctx context.Context, directory domain.EquipmentDirectory) error {
	query := `INSERT INTO equipDirectory (name, releaseYear, type) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, directory.Name, directory.ReleaseYear, directory.TypeId)
	if err != nil {
		return err
	}
	return nil
}

func (r *EquipmentDirectoryRepo) DeleteDirectory(ctx context.Context, id int) error {
	query := `DELETE FROM equipDirectory WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
