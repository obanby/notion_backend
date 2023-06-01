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

func (ds *DocumentService) CreateDocument(parentId string, doc *documents.Document) error {
	isValidParentId := ds.ValidateParentId(parentId)
	err := isValidParentId()
	if err != nil {
		return err
	}
	isParentAndDocUnique := ds.ValidateParentIsNotEqualToChild(parentId, string(doc.Id))
	err = isParentAndDocUnique()
	if err != nil {
		return err
	}
	return ds.Repository.SaveDocument(parentId, doc)
}

type Validator func() error

func (ds *DocumentService) ValidateParentId(parentDocId string) Validator {
	return func() error {
		_, err := ds.Repository.GetDocumentById(parentDocId)
		return err
	}
}

func (ds *DocumentService) ValidateParentIsNotEqualToChild(parentDocId string, documentId string) Validator {
	return func() error {
		if parentDocId == documentId {
			return fmt.Errorf("parent document can't be equal to the child document")
		}
		return nil
	}
}
