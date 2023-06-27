//go:build integration

package repository_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/obanby/notion_backend/boundedContext/documents"
	"github.com/obanby/notion_backend/boundedContext/documents/repository"
	"testing"
)

// BAD TEST, we will refactor it!
func TestNeo4jRepo_New(t *testing.T) {
	want := &repository.Neo4jRepo{}
	got := repository.NewNeo4jRepository()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestNeo4jRepo_SaveDocument(t *testing.T) {
	repo := repository.NewNeo4jRepository()
	doc, err := documents.NewDocument()
	doc.Name = "MyDocument"
	if err != nil {
		t.Fatal(err)
	}
	err = repo.SaveDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
}
