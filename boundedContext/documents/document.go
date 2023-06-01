package documents

import "github.com/google/uuid"

type ID string

type Document struct {
	Id        ID
	Name      string
	Documents []ID // TODO: change to linked list
}

type Option func(document *Document) error

func NewDocument(options ...Option) (*Document, error) {
	doc := &Document{
		Id:        ID(uuid.New().String()),
		Name:      "",
		Documents: nil,
	}
	for _, option := range options {
		err := option(doc)
		if err != nil {
			return &Document{}, err
		}
	}
	return doc, nil
}

func WithDocumentName(name string) Option {
	return func(document *Document) error {
		document.Name = name
		return nil
	}
}
