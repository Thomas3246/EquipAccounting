package database

import (
	"context"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
)

type DocumentRepo interface {
	GetDocumentTypes(context.Context) ([]domain.DocumentType, error)
	AddDocument(context.Context, domain.Document) error
	GetDocumentsView(context.Context, int) ([]domain.DocumentView, error)
	GetDocument(context.Context, int) (domain.Document, error)
	DeleteDocument(ctx context.Context, id int) error
}
