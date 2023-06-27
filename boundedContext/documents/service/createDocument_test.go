package service_test

import (
	"github.com/obanby/notion_backend/boundedContext/documents"
	"github.com/obanby/notion_backend/boundedContext/documents/repository"
	"github.com/obanby/notion_backend/boundedContext/documents/service"
	"testing"
)

func TestDocumentService_CreateDocument(t *testing.T) {
	inMemRepo := repository.NewInMemoryDocumentRepo()
	doc, err := documents.NewDocument()
	if err != nil {
		t.Fatal(err)
	}
	docService, err := service.New(
		service.WithRepository(inMemRepo),
	)
	err = docService.CreateDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	_, err = inMemRepo.GetDocumentById(string(doc.Id))
	if err != nil {
		t.Fatal(err)
	}
}
