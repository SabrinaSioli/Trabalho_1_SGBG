package page

import (
	"Trabalho_1_SGBD/document"
)

type Page interface {
	GetDocuments() []*document.Document
}

type PageImpl struct {
	Header    [5]int //flag binário para indicar se o byte está ocupado 0->livre e 1->ocupado
	PageId    int //identificador da página
	Documents []*document.Document //lista de documentos da página
}

//Construtor nova página
func NewPage(pageId int) *PageImpl {
	return &PageImpl{
		PageId: pageId,
	}
}

//Retorna lista de todos os documentos da página
func (pag *PageImpl) GetDocuments() []*document.Document {
	var documents []*document.Document
	for _, p := range pag.Documents {
		documents = append(documents, p)
	}
	return documents
}

//****INSERIR****

//Verifica se há espaço na página para inserir um novo documento
func (pag *PageImpl) HasSpaceForDocument(documentSize int) bool {
	pageSize := 0
	for _, d := range pag.Documents { //calcula tamanho ocupado pela página
		pageSize += d.DID.Size
	}
	if pageSize+documentSize > 5 { //verifica se o tamanho atual da página + tamanho do documento excede 5
		return false
	}
	return true
}

//Atualiza Header da página após inserir um documento
func (pag *PageImpl) UpdateHeaderOnInsert(size int) {
	for i := 0; i < 5; i++ {
		if pag.Header[i] == 0 && size > 0 {
			pag.Header[i] = 1
			size -= 1
		}
	}
}

//****DELETAR****

//Deleta documento da lista de documentos da página e ajusta possíveis "buracos"
func (pag *PageImpl) DeleteDocument(doc *document.Document)  {
	seqAnt := doc.DID.Seq
	seqProx := 0
	newDocuments := []*document.Document{}

	for _, d := range pag.Documents {
		if d.DID.Seq < doc.DID.Seq { //se o documento vem antes do deletado, não muda seu int seq
			newDocuments = append(newDocuments, d)
		} else if d.DID.Seq > doc.DID.Seq { //se o documento vinha após o deletado, atualizar seu int seq
			seqProx = d.DID.Seq
			d.DID.Seq = seqAnt
			seqAnt = seqProx
			newDocuments = append(newDocuments, d)
		}
	}
	pag.Documents = newDocuments //atualiza lista de documentos da página
}

//Atualiza Header da página após deletar um documento
func (pag *PageImpl) UpdateHeaderOnDelete(doc *document.Document) {
	pageSize := 0
	for _, d := range pag.Documents {
		pageSize += d.DID.Size
	}
	
	pag.Header = [5]int{0,0,0,0,0}

	for i:=0; i < pageSize; i++ {
		pag.Header[i] = 1
	}

}

