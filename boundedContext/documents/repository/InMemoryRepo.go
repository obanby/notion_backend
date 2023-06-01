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

func (i InMemoryDocumentRpeo) SaveDocument(parentId string, doc *documents.Document) error {
	parentDoc := i.inodeTable[parentId]
	parentDoc.Documents = append(parentDoc.Documents, doc.Id)
	i.inodeTable[string(doc.Id)] = doc
	return nil
}

func NewInMemoryDocumentRepo() documents.DocumentRepository {
	return &InMemoryDocumentRpeo{
		inodeTable: map[string]*documents.Document{
			"root": &documents.Document{Id: "root", Name: "", Documents: nil},
		},
	}
}
