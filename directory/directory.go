package directory

import (
	"Trabalho_1_SGBD/document"
	"Trabalho_1_SGBD/page"
	"fmt"
)

type Directory struct {
	UsedPages  []*page.PageImpl //lista de páginas com documentos
	BlankPages []*page.PageImpl //lista de páginas sem documentos (vazias)
}

// Inicializador do diretório
func Init() *Directory {
	blankPages := []*page.PageImpl{}

	for i := 0; i < 20; i++ {
		blankPages = append(blankPages, page.NewPage(i))
	}

	return &Directory{
		BlankPages: blankPages, //somente páginas em branco no início
	}
}

// Retorna todos os documentos do diretório com <content>
func (dir *Directory) Scan() []*document.Document {
	var documents []*document.Document
	for _, pag := range dir.UsedPages {
		documents = append(documents, pag.GetDocuments()...)
	}
	PrintDocuments(documents)

	return documents
}

// Retorna todas as páginas usadas do diretório
func (dir *Directory) ScanPages() []*document.Document {
	var documents []*document.Document
	for _, pag := range dir.UsedPages {
		documents = append(documents, pag.GetDocuments()...)
	}
	PrintPages(documents)

	return documents
}

// Função auxiliar de Scan para imprimir documentos
func PrintDocuments(documents []*document.Document) {
	if len(documents) != 0 {
		for _, doc := range documents {
			//IMPRIMINDO
			fmt.Printf("[ \"%s\" ] ", doc.Content)
		}

		fmt.Println()
	} else {
		fmt.Print("Todas as páginas estão vazias\n\n")
	}
	
}

// Função auxiliar de ScanPages para imprimir páginas usadas
func PrintPages(documents []*document.Document) {
	if len(documents) != 0 {
		idLine := documents[0].DID.PageId //identificador de página para usar na quebra de linha
		for _, doc := range documents {
			if doc.DID.PageId != idLine {
				fmt.Println("")
				idLine = doc.DID.PageId
			}
			//IMPRIMINDO
			fmt.Printf("[ PageId: %d, Seq: %d, Content: \"%s\" ] ", doc.DID.PageId, doc.DID.Seq, doc.Content)
		}
		fmt.Println()
	} else {
		fmt.Print("Todas as páginas estão vazias\n\n")
	}
}

// Retorna DID do documento encontrado ou mensagem de erro se não existir
func (dir *Directory) Seek(content string) {
	foundDoc := false
	fmt.Println("\n___________________________ SEEK ___________________________")
	for _, pag := range dir.UsedPages { //busca nas páginas usadas
		for _, doc := range pag.Documents {
			if doc.Content == content && !foundDoc {
				foundDoc = true //flag para garantir apenas 1ª ocorrência

				//IMPRIMINDO
				fmt.Printf("O documento com \"%s\" foi encontrado!\n", content)
				fmt.Printf("<pageIde: %d, seq: %d, tam: %d> \n\n", doc.DID.PageId, doc.DID.Seq, doc.DID.Size)

				break
			}
		}
	}
	if !foundDoc { //flag indicando página não encontrada
		//IMPRIMINDO
		fmt.Printf("O documento com \"%s\" não foi encontrado!\n\n", content)
	}
}

// Insere documento no diretório se houver espaço
func (dir *Directory) Insert(content string) {
	//fmt.Printf("\n ********** INSERT ********** \n")
	foundPage := false

	//primeiro busca espaço nas páginas já usadas
	for _, pag := range dir.UsedPages {
		if pag.HasSpaceForDocument(len(content)) && !foundPage { //se houver espaço na página
			for index, flag := range pag.Header {
				if flag == 0 { //quando encontrar a primeira posição vazia
					pag.UpdateHeaderOnInsert(len(content))
					newDocument := document.NewDocument(pag.PageId, index, len(content), content)
					pag.Documents = append(pag.Documents, newDocument)
					foundPage = true //flag para garantir apenas uma inserção

					//IMPRIMINDO
					fmt.Printf("\nDocumento com \"%s\" inserido com sucesso! <PageId: %d, Seq: %d, Size: %d> \n\n", content, newDocument.DID.PageId, newDocument.DID.Seq, newDocument.DID.Size)

					break
				}
			}
		}
	}

	//caso não encontre espaço nas páginas usadas, usa as páginas em branco
	if !foundPage { //caso não tenha espaço nas páginas usadas, buscar nas em branco
		if len(dir.BlankPages) > 0 { //se houver página em branco
			newPage := &page.PageImpl{}
			newPage, dir.BlankPages = dir.BlankPages[0], dir.BlankPages[1:] //removendo página da lista em branco
			newPage.UpdateHeaderOnInsert(len(content))
			newDocument := document.NewDocument(newPage.PageId, 0, len(content), content) //inserindo documento na página
			newPage.Documents = append(newPage.Documents, newDocument)
			dir.UsedPages = append(dir.UsedPages, newPage) //adicionando página na lista das usadas

			//IMPRIMINDO
			fmt.Printf("\nDocumento com \"%s\" inserido com sucesso! <PageId: %d, Seq: %d, Size: %d> \n\n", content, newDocument.DID.PageId, newDocument.DID.Seq, newDocument.DID.Size)

		} else { //não há espaço no diretório para inserir o documento
			//IMPRIMINDO
			fmt.Printf("\nNão há espaço suficiente nas páginas para inserir \"%s\"!\n\n", content)
		}
	}

	dir.sortPages() //ordena páginas usando pageId
}

// Deleta primeiro documento encontrado com o conteúdo ou exibe mensagem de erro se não existir
func (dir *Directory) Delete(content string) {
	foundDoc := false
	for index, pag := range dir.UsedPages { //busca nas páginas usadas
		if !foundDoc {
			for _, doc := range pag.Documents { //varrendo documentos da página
				if doc.Content == content && !foundDoc {
					pag.DeleteDocument(doc)       //deletando documento do array da página
					pag.UpdateHeaderOnDelete(doc) //atualizando header com flag de cada byte usado/branco

					//se a página ficar em branco após deleção
					if pag.Header[0] == 0 {
						dir.UsedPages = append(dir.UsedPages[:index], dir.UsedPages[index+1:]...) //remove página da lista de usadas
						dir.BlankPages = append(dir.BlankPages, pag)                              //adiciona página na lista de páginas em branco
					}
					foundDoc = true

					//IMPRIMINDO
					fmt.Printf("\nDocumento com \"%s\" deletado com sucesso!\n", content)
					fmt.Print("Página resultante: ")
					fmt.Printf("PageId:%d, Header:[%d, %d, %d, %d, %d], Documents: {", pag.PageId, pag.Header[0], pag.Header[1], pag.Header[2], pag.Header[3], pag.Header[4])
					fmt.Print(pag.Documents)
					fmt.Printf("}\n\n")
				}
			}
		}

		dir.sortPages() //ordena páginas usando pageId
	}

	if !foundDoc { //flag indicando página não encontrada
		fmt.Printf("\nDocumento com \"%s\" não foi encontrado para deletar.\n\n", content)
	}
}

// Ordena listas UsedPages e BlankPages usando BubbleSort, com pageId como parâmetro
func (dir *Directory) sortPages() {

	//UsedPages
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(dir.UsedPages)-1 {
			if dir.UsedPages[i].PageId > dir.UsedPages[i+1].PageId {
				dir.UsedPages[i], dir.UsedPages[i+1] = dir.UsedPages[i+1], dir.UsedPages[i]
				isDone = false
			}
			i++
		}
	}

	//BlankPages
	isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(dir.BlankPages)-1 {
			if dir.BlankPages[i].PageId > dir.BlankPages[i+1].PageId {
				dir.BlankPages[i], dir.BlankPages[i+1] = dir.BlankPages[i+1], dir.BlankPages[i]
				isDone = false
			}
			i++
		}
	}
}
