package repository

import (
	"fmt"
	"github.com/obanby/notion_backend/boundedContext/documents"
)

type InMemoryDocumentRpeo struct {
	inodeTable map[string]*documents.Document
}

func (i InMemoryDocumentRpeo) GetDocumentById(id string) (*documents.Document, error) {
	if document, exists := i.inodeTable[id]; exists {
		return document, nil
	}
	return nil, fmt.Errorf("document with id %v doesn't exist", id)
}

func (i InMemoryDocumentRpeo) SaveDocument(doc *documents.Document) error {
	i.inodeTable[string(doc.Id)] = doc
	return nil
}

func NewInMemoryDocumentRepo() documents.DocumentRepository {
	return &InMemoryDocumentRpeo{
		inodeTable: map[string]*documents.Document{},
	}
}
