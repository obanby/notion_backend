package documents

type DocumentRepository interface {
	GetDocumentById(id string) (*Document, error)
	SaveDocument(document *Document) error
}
