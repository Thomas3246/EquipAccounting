package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type DocumentRepo struct {
	db *sql.DB
}

func NewDocumentRepo(db *sql.DB) *DocumentRepo {
	return &DocumentRepo{db: db}
}

func (r *DocumentRepo) GetDocumentTypes(ctx context.Context) (docTypes []domain.DocumentType, err error) {
	query := `SELECT * FROM docType`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := domain.DocumentType{}
		err = rows.Scan(&t.Id, &t.Name)
		if err != nil {
			return nil, err
		}
		docTypes = append(docTypes, t)
	}

	return docTypes, nil
}

func (r *DocumentRepo) AddDocument(ctx context.Context, document domain.Document) error {
	query := `INSERT INTO documents (request, type, file, addDate, user, name) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, document.RequestId, document.Type, document.File, document.AddDate, document.UserId, document.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *DocumentRepo) GetDocumentsView(ctx context.Context, requestId int) (docs []domain.DocumentView, err error) {
	query := `SELECT d.id, d.request, t.name, d.addDate, u.login, d.name 
			  FROM documents AS d
			  INNER JOIN docType AS t ON d.type = t.id
			  INNER JOIN users AS u ON d.user = u.id 
			  WHERE request = ?`

	rows, err := r.db.QueryContext(ctx, query, requestId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d := domain.DocumentView{}
		err = rows.Scan(&d.Id, &d.RequestId, &d.Type, &d.AddDate, &d.UserLogin, &d.Name)
		if err != nil {
			return nil, err
		}
		docs = append(docs, d)
	}
	return docs, nil
}

func (r *DocumentRepo) GetDocument(ctx context.Context, id int) (doc domain.Document, err error) {
	query := `SELECT * FROM documents WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return domain.Document{}, row.Err()
	}

	err = row.Scan(&doc.Id, &doc.RequestId, &doc.Type, &doc.File, &doc.AddDate, &doc.UserId, &doc.Name)
	if err != nil {
		return domain.Document{}, err
	}
	return doc, nil
}

func (r *DocumentRepo) DeleteDocument(ctx context.Context, id int) error {
	query := `DELETE FROM documents WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
