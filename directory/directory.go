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

//Inicializador do diretório
func Init(blankPages []*page.PageImpl) *Directory {
	return &Directory{
		BlankPages: blankPages, //somente páginas em branco no início
	}
}

//Retorna todos os documentos do diretório <content>
func (dir *Directory) Scan() []*document.Document {
	var documents []*document.Document
	for _, pag := range dir.UsedPages {
		documents = append(documents, pag.GetDocuments()...)
	}
	PrintDocuments(documents)

	return documents
}

//Retorna todas as páginas usadas do diretório 
func (dir *Directory) ScanPages() []*document.Document {
	var documents []*document.Document
	for _, pag := range dir.UsedPages {
		documents = append(documents, pag.GetDocuments()...)
	}
	PrintPages(documents)

	return documents
}

//Função auxiliar de Scan para imprimir documentos
func PrintDocuments(documents []*document.Document) {
	fmt.Println(" *************** EXECUTANDO SCAN ***************")
	for _, doc := range documents {
		fmt.Printf("[ \"%s\" ] ", doc.Content)
	}
	fmt.Println()
}

//Função auxiliar de ScanPages para imprimir páginas usadas
func PrintPages(documents []*document.Document) {
	fmt.Println(" ********** EXECUTANDO SCAN WITH PAGES **********")
	idLine := documents[0].DID.PageId
	for _, doc := range documents {
		if doc.DID.PageId != idLine {
			fmt.Println("")
			idLine = doc.DID.PageId
		}
		fmt.Printf("[ PageId: %d, Seq: %d, Content: \"%s\" ] ", doc.DID.PageId, doc.DID.Seq, doc.Content)
	}
	fmt.Println()
}

//Retorna DID do documento encontrado ou mensagem de erro se não existir
func (dir *Directory) Seek(content string) {
	find := false
	fmt.Println("\n******************* SEEK  *******************")
	for _, pag := range dir.UsedPages { //busca nas páginas usadas
		for _, doc := range pag.Documents {
			if doc.Content == content && !find {
				fmt.Printf("\nO documento com \"%s\" foi encontrado!\n", content)
				fmt.Printf("<pageIde: %d, seq: %d, tam: %d> \n\n", doc.DID.PageId, doc.DID.Seq, doc.DID.Size)
				find = true
			}
		}
	}
	if !find { //flag indicando página não encontrada
		fmt.Printf("\nO documento com \"%s\" não foi encontrado!\n", content)
	}
}


//Insere documento no diretório se houver espaço
func (dir *Directory) Insert(content string) {
	fmt.Printf("\n ********** INSERT ********** \n")
	if len(content) > 0 && len(content) <= 5 { //verificando se o conteúdo está no padrão
		foundPage := false
		for _, pag := range dir.UsedPages { //primeiro busca espaço nas páginas usadas
			if foundPage {
				break
			}
			if pag.HasSpaceForDocument(len(content)) { //se houver espaço na página
				for i, flag := range pag.Header {
					if flag == 0 { //quando encontrar a primeira posição vazia 
						pag.UpdateHeaderOnInsert(len(content))
						newDocument := document.NewDocument(pag.PageId, i, len(content), content)
						pag.Documents = append(pag.Documents, newDocument)
						fmt.Printf("\nDocumento com \"%s\" inserido com sucesso! <PageId: %d, Seq: %d, Size: %d> \n\n", content, newDocument.DID.PageId, newDocument.DID.Seq, newDocument.DID.Size)
						foundPage = true
						break
					}
				}
			}
		}
		if !foundPage { //caso não tenha espaço nas páginas usadas, buscar nas em branco
			if len(dir.BlankPages) > 0 { //se houver página em branco
				dir.sortPages()
				newPage := &page.PageImpl{}
				newPage, dir.BlankPages = dir.BlankPages[0], dir.BlankPages[1:]
				newPage.UpdateHeaderOnInsert(len(content))
				newDocument := document.NewDocument(newPage.PageId, 0, len(content), content)
				newPage.Documents = append(newPage.Documents, newDocument)
				dir.UsedPages = append(dir.UsedPages, newPage)
				fmt.Printf("\nDocumento com \"%s\" inserido com sucesso! <PageId: %d, Seq: %d, Size: %d> \n\n", content, newDocument.DID.PageId, newDocument.DID.Seq, newDocument.DID.Size)
			} else { //não há espaço no diretório para inserir o documento
				fmt.Printf("\nNão há espaço suficiente nas páginas para inserir \"%s\"!\n\n", content)
			}
		}
	}
	dir.sortPages() //ordena páginas usando pageId
}

//Deleta primeiro documento encontrado com o conteúdo ou exibe mensagem de erro se não existir
func (dir *Directory) Delete(content string) {
	fmt.Println("\n******************* DELETE  *******************")
	find := false
	for index, pag := range dir.UsedPages { //busca nas páginas usadas
		if !find {
			for _, doc := range pag.Documents {
				if doc.Content == content && !find {
					find = true
					pag.DeleteDocument(doc) //deletando documento do array da página
					pag.UpdateHeaderOnDelete(doc) //atualizando header com flag de cada byte usado/branco
					if pag.Header[0] == 0 { //se a página ficar em branco após deleção
						dir.UsedPages = append(dir.UsedPages[:index], dir.UsedPages[index+1:]...) //remove página da lista de usadas
						dir.BlankPages = append(dir.BlankPages, pag) //adiciona página na lista de páginas em branco
					}
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

	if !find { //flag indicando página não encontrada
		fmt.Printf("\nDocumento com \"%s\" não foi encontrado para deletar.\n\n", content)
	}
}

//Ordena listas UsedPages e BlankPages usando BubbleSort, com pageId como parâmetro
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


