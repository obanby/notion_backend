package service

import (
	"fmt"
	"github.com/obanby/notion_backend/boundedContext/documents"
)

type DocumentService struct {
	Repository documents.DocumentRepository
}

type Option func(service *DocumentService) error

func WithRepository(repository documents.DocumentRepository) Option {
	return func(service *DocumentService) error {
		if repository == nil {
			return fmt.Errorf("repository must not be nil")
		}
		service.Repository = repository
		return nil
	}
}

func New(options ...Option) (*DocumentService, error) {
	document := &DocumentService{}
	for _, option := range options {
		err := option(document)
		if err != nil {
			return nil, err
		}
	}
	return document, nil
}

/* TODO:
2 - Implement DTO if needed
3 - Implmenet DB for create document
*/

func (ds *DocumentService) CreateDocument(doc *documents.Document) error {
	return ds.Repository.SaveDocument(doc)
}

type Validator func() error
