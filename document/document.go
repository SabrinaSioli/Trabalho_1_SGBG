package document

//DID e conteúdo como atributos de Documento
type Document struct {
	DID  DID
	Content string
}

//Tupla <int page id, int seq, int tam> 
type DID struct {
	PageId int //identificador página
	Seq    int //endereço da posição inicial do documento na página
	Size   int //tamanho em bytes do documento
}

//Construtor de novo documento
func NewDocument(pageId int, seq int, size int, content string) *Document {
	did := DID{
		PageId: pageId,
		Seq:    seq,
		Size:   size,
	}
	return &Document{
		DID:  did,
		Content: content,
	}
}
