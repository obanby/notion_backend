package service_test

import (
	"github.com/obanby/notion_backend/boundedContext/documents"
	"github.com/obanby/notion_backend/boundedContext/documents/repository"
	"github.com/obanby/notion_backend/boundedContext/documents/service"
	"testing"
)

func TestDocumentService_CreateDocument(t *testing.T) {
	parentDoc, err := documents.NewDocument()
	inMemRepo := repository.NewInMemoryDocumentRepo()
	err = inMemRepo.SaveDocument("root", parentDoc)
	if err != nil {
		t.Fatal(err)
	}
	testcases := []struct {
		parentId    string
		shouldError bool
		doc         *documents.Document
	}{
		{
			parentId:    "root",
			shouldError: false,
			doc: &documents.Document{
				Id:        "1234",
				Documents: []documents.ID{},
			},
		},
		{
			parentId:    string(parentDoc.Id),
			shouldError: false,
			doc: &documents.Document{
				Id:        "2345",
				Documents: []documents.ID{},
			},
		},
		{
			parentId:    "",
			shouldError: true,
			doc: &documents.Document{
				Id:        "6789",
				Documents: []documents.ID{},
			},
		},
		{
			parentId:    "invalid",
			shouldError: true,
			doc: &documents.Document{
				Id:        "10111",
				Documents: []documents.ID{},
			},
		},
		{
			parentId:    string(parentDoc.Id),
			shouldError: true,
			doc:         parentDoc,
		},
	}

	docService, err := service.New(
		service.WithRepository(inMemRepo),
	)
	if err != nil {
		t.Fatal(err)
	}
	for idx, tc := range testcases {
		err = docService.CreateDocument(tc.parentId, tc.doc)
		if tc.shouldError && err == nil {
			t.Fatalf("expected testcase: %v to error, but it didn't.", idx)
		}
		if tc.shouldError {
			continue
		}
		_, err = inMemRepo.GetDocumentById(string(tc.doc.Id))
		if err != nil {
			t.Fatal(err)
		}
	}
}
