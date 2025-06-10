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

func (r *RequestRepo) GetAllClosedDetail(ctx context.Context) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name, description, users.name, requestStatus.name, request.createdAt, COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus on request.status = requestStatus.id
			  INNER JOIN equipment on request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  WHERE request.status = 2`

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
	query := `SELECT request.id, requestType.name, description, users.name, requestStatus.name, request.createdAt, COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name
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

func (r *RequestRepo) GetAllClosedForUserDetail(ctx context.Context, login string) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name, description, users.name, requestStatus.name, request.createdAt, COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus on request.status = requestStatus.id
			  INNER JOIN equipment on request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  INNER JOIN users AS requester ON requester.login = ?
			  INNER JOIN department ON requester.department = department.id
			  WHERE 
			  request.status = 2 AND users.department = department.id`

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

func (r *RequestRepo) GetAllUserActiveDetail(ctx context.Context, login string) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name AS type_name, request.description, users.name AS author_name, requestStatus.name AS status_name, request.createdAt, 
			  COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name AS equipment_info
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus ON request.status = requestStatus.id
			  INNER JOIN equipment ON request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  WHERE request.status = 1 AND users.login = ?`

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

func (r *RequestRepo) GetAllUserClosedDetail(ctx context.Context, login string) (requests []domain.RequestView, err error) {
	query := `SELECT request.id, requestType.name AS type_name, request.description, users.name AS author_name, requestStatus.name AS status_name, request.createdAt, 
			  COALESCE(request.closedAt, '') as closedAt, equipment.invNum || ' - ' || equipDirectory.name AS equipment_info
			  FROM request
			  INNER JOIN requestType ON request.requestType = requestType.id
			  INNER JOIN users ON request.requestAuthor = users.id
			  INNER JOIN requestStatus ON request.status = requestStatus.id
			  INNER JOIN equipment ON request.equipment = equipment.id
			  INNER JOIN equipDirectory ON equipment.directory = equipDirectory.id
			  WHERE request.status = 2 AND users.login = ?`

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

func (r *RequestRepo) GetRequestTypes(ctx context.Context) (types []domain.RequestType, err error) {
	query := `SELECT id, name FROM requestType`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := domain.RequestType{}
		err = rows.Scan(&t.Id, &t.Name)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	return types, nil
}

func (r *RequestRepo) AddRequest(ctx context.Context, request domain.Request) error {
	query := `INSERT INTO request (requestType, description, requestAuthor, status, createdAt, equipment) 
			  VALUES (?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, request.Type, request.Description, request.Author, request.Status, request.CreatedAt, request.Equipment)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestRepo) GetRequestById(ctx context.Context, id int) (request domain.Request, err error) {
	query := `SELECT id, requestType, description, requestAuthor, status, createdAt, COALESCE(request.closedAt, '') as closedAt, equipment, COALESCE(result, 0) as result, COALESCE(resultDescr, '') as resultDescr 
			  FROM request WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return request, err
	}

	err = row.Scan(&request.Id, &request.Type, &request.Description, &request.Author, &request.Status, &request.CreatedAt, &request.ClosedAt, &request.Equipment, &request.Result, &request.ResultDescr)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (s *RequestRepo) UpdateDescription(ctx context.Context, reqId int, descr string) error {
	query := `UPDATE request 
			  SET description = ?
			  WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, descr, reqId)
	return err
}

func (s *RequestRepo) UpdateRequest(ctx context.Context, request domain.Request) error {
	query := `UPDATE request
			  SET requestType = ?, description = ?, equipment = ?
			  WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, request.Type, request.Description, request.Equipment, request.Id)
	return err
}

func (r *RequestRepo) GetResults(ctx context.Context) (results []domain.RequestResult, err error) {
	query := `SELECT id, name FROM requestResult`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := domain.RequestResult{}
		err = rows.Scan(&r.Id, &r.Name)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func (r *RequestRepo) GetRequestsWithEquipment(ctx context.Context, equipId int) (requests []domain.Request, err error) {
	query := `SELECT id, requestType, description, requestAuthor, status, createdAt, COALESCE(closedAt, '') as closedAt, equipment, COALESCE(result, 0) as result, COALESCE(resultDescr, '') as resultDescr
			  FROM request
			  WHERE status = 1 AND equipment = ?`

	rows, err := r.db.QueryContext(ctx, query, equipId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := domain.Request{}
		err = rows.Scan(&r.Id, &r.Type, &r.Description, &r.Author, &r.Author, &r.CreatedAt, &r.ClosedAt, &r.Equipment, &r.Result, &r.ResultDescr)
		if err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}
	return requests, nil
}

func (r *RequestRepo) CloseRequest(ctx context.Context, requestId int, resultId int, closedAt string, resultDescr string) error {
	query := `UPDATE request
			  SET status = 2, closedAt = ?, result = ?, resultDescr = ?
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, closedAt, resultId, resultDescr, requestId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestRepo) GetReportData(ctx context.Context, requestId int, userAdminLogin string) (report domain.RequestReport, err error) {
	query := `  SELECT 
				  	r.id, r.requestType, COALESCE(r.description, '') as description, r.createdAt, r.result, r.resultDescr, 
				  	(SELECT u.name FROM users AS u WHERE u.login = ?),
				  	e.invNum, e.purchDate, e.regDate,
				  	ed.name || ' ' || ed.releaseYear,
				  	dep.name || ' ' || depdiv.name
			  	FROM
					request AS r 
					INNER JOIN equipment AS e ON r.equipment = e.id
					INNER JOIN equipDirectory AS ed ON e.directory = ed.id
					INNER JOIN department AS dep ON e.department = dep.id
					INNER JOIN departmentDivisions AS depdiv ON dep.division = depdiv.id	
				WHERE r.id = ?`

	row := r.db.QueryRowContext(ctx, query, userAdminLogin, requestId)
	if row.Err() != nil {
		return domain.RequestReport{}, row.Err()
	}
	err = row.Scan(&report.RequestId, &report.TypeId, &report.Desctiption, &report.CreatedAt, &report.ResultId, &report.ResultDescr, &report.AuthorName, &report.Equipment.InvNum,
		&report.Equipment.PurchDate, &report.Equipment.RegDate, &report.Equipment.Directory, &report.Equipment.Department)
	if err != nil {
		return domain.RequestReport{}, err
	}
	return report, nil
}
