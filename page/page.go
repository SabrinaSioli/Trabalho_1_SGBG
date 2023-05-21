package page

import (
	"Trabalho_1_SGBD/document"
)

const MaxPageSize = 5

type Page interface {
	GetDocuments() []*document.Document
} 

type PageImpl struct {
	Header [5]int
	PageId int
	Documents []*document.Document
}

func NewPage(pageId int) *PageImpl {
	return &PageImpl{
		PageId: pageId,
	}
}

/*
func (pag *PageImpl) IsBlankedOnDelete(doc *document.Document) bool {
	isBlanked := false
	newPage := &PageImpl{}

	return isBlanked
}
*/

func (pag *PageImpl) GetDocuments() []*document.Document {
	var documents []*document.Document
	for _, p := range pag.Documents {
		documents = append(documents, p)
	}
	return documents
}

func (pag *PageImpl) HaveSpaceForDocument(documentSize int) bool  {
	pageSize := 0
	for _, d := range pag.Documents {
		pageSize += d.DID.Tam
	}
	if pageSize + documentSize < 5 {
		return false
	}
	return true
}

func (pag *PageImpl) FindDocumentByContent(content string) (*document.Document, bool) {
	for _, doc := range pag.Documents {
		if (doc.Info == content) {
			return doc, true
		}
	}
	return &document.Document{}, false
}


