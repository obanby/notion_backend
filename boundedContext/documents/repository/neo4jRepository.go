package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/obanby/notion_backend/boundedContext/documents"
	"sync"
)

type Neo4jRepo struct {
	uri      string
	username string
	password string
	driver   neo4j.DriverWithContext
}

var lock = sync.Mutex{}
var instance *Neo4jRepo

func newNeo4jRepository() (*Neo4jRepo, error) {
	repo := &Neo4jRepo{
		uri:      "neo4j://localhost:7687",
		username: "neo4j",
		password: "Password1",
	}
	driver, err := neo4j.NewDriverWithContext(
		repo.uri,
		neo4j.BasicAuth(repo.username, repo.password, ""),
	)
	if err != nil {
		return nil, err
	}
	repo.driver = driver
	return repo, nil
}

func GetNeo4jInstance() (*Neo4jRepo, error) {
	lock.Lock()
	defer lock.Unlock()
	if instance != nil {
		return instance, nil
	}
	repo, err := newNeo4jRepository()
	instance = repo
	return instance, err
}

func (n *Neo4jRepo) GetDocumentById(id string) (*documents.Document, error) {
	ctx := context.TODO()
	session := n.driver.NewSession(ctx, neo4j.SessionConfig{})
	doc, err := neo4j.ExecuteRead(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) (*documents.Document, error) {
			cypher := `
				MATCH (d:Document)
				WHERE d.id = $id
				RETURN d;
			`

			doc := documents.Document{}
			result, err := tx.Run(ctx, cypher, map[string]any{
				"id": id,
			})
			if err != nil {
				return &doc, err
			}

			record, err := result.Single(ctx)
			if err != nil {
				return &doc, err
			}

			node, isFound := record.Get("d")
			if !isFound {
				return &doc, fmt.Errorf("document node is not found in database")
			}

			docNode, isNode := node.(neo4j.Node)
			if !isNode {
				return &doc, fmt.Errorf("document is not a node")
			}

			docName, err := neo4j.GetProperty[string](docNode, "name")
			if err != nil {
				return &doc, fmt.Errorf("property name is not found")
			}

			return &documents.Document{
				Id:   documents.ID(id),
				Name: docName,
			}, nil
		},
	)
	return doc, err
}

func (n *Neo4jRepo) SaveDocument(document *documents.Document) error {
	ctx := context.TODO()
	session := n.driver.NewSession(ctx, neo4j.SessionConfig{})
	_, err := neo4j.ExecuteWrite(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) (any, error) {
			chyper := `
				MERGE (d:Document)
				set d.id = $documentId
				set d.name = $documentName
				RETURN d
			`
			result, err := tx.Run(ctx, chyper, map[string]any{
				"documentId":   document.Id,
				"documentName": document.Name,
			})
			if err != nil {
				return "", err
			}
			_, err = neo4j.SingleTWithContext(
				ctx,
				result,
				func(record *neo4j.Record) (any, error) {
					_, _, err := neo4j.GetRecordValue[neo4j.Node](record, "d")
					return "", err
				},
			)
			return "", err
		},
	)
	return err
}
