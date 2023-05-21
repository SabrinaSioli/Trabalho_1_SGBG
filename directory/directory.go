package directory

import (
	"Trabalho_1_SGBD/page"
	"Trabalho_1_SGBD/document"
	"fmt"
)

type Directory struct {
	UsedPages []*page.PageImpl
	BlankPages []*page.PageImpl
}

func Init(blankPages []*page.PageImpl) *Directory{
	return &Directory{
		BlankPages: blankPages,
	}
}

func (dir *Directory) Scan() []*document.Document {
	var documents [] *document.Document
	for _, pag := range dir.UsedPages {
		documents = append(documents, pag.GetDocuments()...)
	}
	PrintDocuments(documents)

	return documents
}

func (dir *Directory) Seek(content string) {
	find := false
	fmt.Println("******************* SEEK  *******************")
	for _, pag := range dir.UsedPages {
		for _, doc := range pag.Documents {
			if(doc.Info == content) {
				fmt.Println("O documento foi encontrado!")
				fmt.Printf("<pageIde: %d, seq: %d, tam: %d>", doc.DID.PageId, doc.DID.Seq, doc.DID.Tam)
				find = true
				break
			}
		}
	}
	if !find {
		fmt.Println("O documento não foi encontrado!")
	}
}
/*
func (dir *Directory) Delete(content string) {
	find := false
	for _, pag := range dir.UsedPages {
		for _, doc := range pag.Documents {
			if doc.Info == content {
				pag.IsBlankedOnDelete(doc)
				fmt.Println("Deletado com sucesso")

			}
		}
	}
}
*/


func PrintDocuments(documents []*document.Document) {
	fmt.Println(" *************** EXECUTANDO SCAN ***************")
	for _, doc := range documents {
		fmt.Printf("[ \"%s\" ] ", doc.Info)
	}
	fmt.Println()
}

func (dir *Directory) Insert(content string) {
	foundPage := false
	for _, pag := range dir.UsedPages {
		if foundPage {
			break
		}
		 if pag.HaveSpaceForDocument(len(content)) {
			for i, b := range pag.Header {
				if b == 'f' {
					newDocument := document.NewDocument(pag.PageId, i, len(content), content)
					pag.Documents = append(pag.Documents, newDocument)
					foundPage = true
					break
				}
			}
		 }
	}
	if !foundPage {
		if len(dir.BlankPages) > 1 {
			newPage := &page.PageImpl{}
			newPage, dir.BlankPages = dir.BlankPages[0], dir.BlankPages[1:]
			dir.UsedPages = append(dir.UsedPages, newPage)
		}
	}
}

// TESTES
func PrintDocumentsComPages(documents []*document.Document) {
	fmt.Println(" *************** EXECUTANDO SCAN ***************")
	for i, doc := range documents {
		fmt.Printf(" --- Documento %d ---\n", i)
		fmt.Printf("Página %d: [%s]\n", doc.DID.PageId, doc.Info)
	}
}