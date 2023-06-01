package documents

type IDocumentService interface {
	CreateDocument(document *Document) error
}
