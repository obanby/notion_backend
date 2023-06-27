package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/obanby/notion_backend/boundedContext/documents"
)

type Neo4jRepo struct {
	//
}

func NewNeo4jRepository() *Neo4jRepo {
	return &Neo4jRepo{}
}

func (n *Neo4jRepo) GetDocumentById(id string) (*documents.Document, error) {
	//TODO implement me
	panic("implement me")
}

func (n *Neo4jRepo) SaveDocument(document *documents.Document) error {
	uri := "neo4j://localhost:7687"
	userName := "neo4j"
	password := "Password1"
	if len(userName) == 0 || len(password) == 0 {
		return fmt.Errorf("invalid neo4j username or password")
	}
	driver, err := neo4j.NewDriverWithContext(
		uri,
		neo4j.BasicAuth(userName, password, ""),
	)
	if err != nil {
		return err
	}
	ctx := context.Background()
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	_, err = neo4j.ExecuteWrite(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) (documents.Document, error) {
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
				return documents.Document{}, err
			}
			_, err = neo4j.SingleTWithContext[documents.Document](
				ctx,
				result,
				func(record *neo4j.Record) (documents.Document, error) {
					node, _, _ := neo4j.GetRecordValue[neo4j.Node](record, "d")
					fmt.Println(node.Props["id"], node.Props["name"])
					return documents.Document{}, nil
				},
			)
			if err != nil {
				return documents.Document{}, err
			}
			return documents.Document{
				Id:   document.Id,
				Name: document.Name,
			}, nil
		},
	)
	if err != nil {
		panic(err)
	}

	return nil
	// new session
	// neo4j transaction
	// cypher with arguments
}
