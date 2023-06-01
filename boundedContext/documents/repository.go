package documents

type DocumentRepository interface {
	GetDocumentById(id string) (*Document, error)
	SaveDocument(parentId string, document *Document) error
}
