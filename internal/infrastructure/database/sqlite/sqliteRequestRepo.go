package sqlite

import (
	"context"
	"database/sql"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/request"
)

type RequestRepo struct {
	db *sql.DB
}

func NewRequestRepo(db *sql.DB) *RequestRepo {
	return &RequestRepo{db: db}
}

func (r *RequestRepo) GetAllActive(ctx context.Context) (requests []domain.Request, err error) {
	query := "SELECT request.id, request.workStation, request.requestType, request.description, request.requestAuthor, request.status, request.createdAt FROM request INNER JOIN requestStatus on request.status = requestStatus.id WHERE requestStatus.name = \"active\" "

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
		err = rows.Scan(&r.Id, &r.WorkStation, &r.Type, &r.Description, &r.Author, &r.Status, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}

	return requests, nil
}
