package document

type Document struct {
	DID DID
	Info string
}

type DID struct {
	PageId int
	Seq int
	Tam int
}

func NewDocument(pageId int, seq int, tam int, info string) *Document {
	did := DID {
		PageId: pageId, 
		Seq: seq, 
		Tam: tam,
	}
	return &Document{
		DID: did,
		Info: info,
	}
}
