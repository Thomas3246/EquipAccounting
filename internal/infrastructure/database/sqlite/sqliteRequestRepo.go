package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type RequestRepo struct {
	db *sql.DB
}

func NewRequestRepo(db *sql.DB) *RequestRepo {
	return &RequestRepo{db: db}
}

func (r *RequestRepo) GetAllActive(ctx context.Context) (requests []domain.Request, err error) {
	query := `SELECT id, requestType, description, requestAuthor, status, createdAt, closedAt, equipment 
			  FROM request 
			  WHERE status = 1`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := domain.Request{}
		err = rows.Scan(&r.Id, &r.Type, &r.Description, &r.Author, &r.Status, &r.CreatedAt, &r.ClosedAt, &r.Equipment)
		if err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}

	return requests, nil
}

func (r *RequestRepo) GetAllActiveDetail(ctx context.Context) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name, description, users.name, requestStatus.name, request.createdAt, COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus on request.status = requestStatus.id
			  INNER JOIN equipment on request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  WHERE request.status = 1`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := domain.RequestView{}
		err = rows.Scan(&r.Id, &r.Type, &r.Description, &r.Author, &r.Status, &r.CreatedAt, &r.ClosedAt, &r.Equipment)
		if err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}

	return requests, nil
}

func (r *RequestRepo) GetAllActiveForUserDetail(ctx context.Context, login string) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name, description, users.name, requestStatus.name, request.createdAt, request.closedAt, equipment.invNum || ' - ' || equipDirectory.name
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus on request.status = requestStatus.id
			  INNER JOIN equipment on request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  INNER JOIN users AS requester ON requester.login = ?
			  INNER JOIN department ON requester.department = department.id
			  WHERE 
			  request.status = 1 AND users.department = department.id`

	rows, err := r.db.QueryContext(ctx, query, login)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := domain.RequestView{}
		err = rows.Scan(&r.Id, &r.Type, &r.Description, &r.Author, &r.Status, &r.CreatedAt, &r.ClosedAt, &r.Equipment)
		if err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}

	return requests, nil
}
