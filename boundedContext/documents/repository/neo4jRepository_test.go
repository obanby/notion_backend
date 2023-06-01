//go:build integration

package repository_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/obanby/notion_backend/boundedContext/documents"
	"github.com/obanby/notion_backend/boundedContext/documents/repository"
	"sync"
	"testing"
)

func TestNeo4jRepo_New(t *testing.T) {
	var want *repository.Neo4jRepo
	var got *repository.Neo4jRepo
	var err error
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		want, err = repository.GetNeo4jInstance()
		if err != nil {
			t.Fatal(err)
		}
		wg.Done()
	}()
	go func() {
		got, err = repository.GetNeo4jInstance()
		if err != nil {
			t.Fatal(err)
		}
		wg.Done()
	}()
	wg.Wait()
	if want != got {
		t.Fatalf("Different instances of the repo! want=%p, got=%p", want, got)
	}
}

// TODO: Refactor to two tests using a initialization phase
func TestNeo4jRepo_SavingAndRetrievingDocuments(t *testing.T) {
	repo, err := repository.GetNeo4jInstance()
	if err != nil {
		t.Fatal(err)
	}
	doc, err := documents.NewDocument()
	doc.Name = "MyDocument"
	if err != nil {
		t.Fatal(err)
	}
	err = repo.SaveDocument(doc)
	if err != nil {
		t.Fatal(err)
	}

	want := doc
	got, err := repo.GetDocumentById(string(doc.Id))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
