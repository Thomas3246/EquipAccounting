package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type DocumentService struct {
	repo database.DocumentRepo
}

func NewDocumentService(repo database.DocumentRepo) *DocumentService {
	return &DocumentService{repo: repo}
}

func (s *DocumentService) GetDocumentTypes() ([]domain.DocumentType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docTypes, err := s.repo.GetDocumentTypes(ctx)
	if err != nil {
		return nil, err
	}
	return docTypes, nil
}

func (s *DocumentService) AddDocument(document domain.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	document.AddDate = time.Now().Format("2006.01.02")

	err := s.repo.AddDocument(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func (s *DocumentService) GetDocumentsViewForRequest(requestId int) ([]domain.DocumentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docs, err := s.repo.GetDocumentsView(ctx, requestId)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func (s *DocumentService) GetDocument(id int) (domain.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc, err := s.repo.GetDocument(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Document{}, ErrNotFound
		}
		return domain.Document{}, err
	}
	return doc, nil
}

func (s *DocumentService) DeleteDocument(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.DeleteDocument(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}
	return nil
}
