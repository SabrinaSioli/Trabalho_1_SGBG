package main

import (
	"Trabalho_1_SGBD/directory"
	"Trabalho_1_SGBD/document"
	"Trabalho_1_SGBD/page"
	"fmt"
)

const MaxPage = 20;

func main() {
	//Testes Iniciais
			doc := document.NewDocument(0,0,0,"str")
			fmt.Println(doc)
			p := page.PageImpl{Header: [5]int{1, 2, 3, 4, 5}, Documents: []*document.Document{doc}}
			fmt.Println(p)
			d := directory.Directory{UsedPages: []*page.PageImpl{&p}, BlankPages: nil}
			fmt.Println(d.UsedPages)

			d1 := document.NewDocument(0, 0, 2, "d1")
			d2 := document.NewDocument(0, 2, 2, "d2")
			d3 := document.NewDocument(1, 0, 2, "d3")
			d4 := document.NewDocument(1, 2, 2, "d4")
			p1 := page.PageImpl{Header: [5]int{1, 2, 3, 4, 5}, Documents: []*document.Document{d1, d2}}
			p2 := page.PageImpl{Header: [5]int{1, 2, 3, 4, 5}, Documents: []*document.Document{d3, d4}}
			p3 := page.PageImpl{Header: [5]int{1, 2, 3, 4, 5}, Documents: []*document.Document{}}
			usedPages := []*page.PageImpl{&p1,&p2}
			//blankPages := []*page.PageImpl{&p3}
	

		// Insert Pages
		blankPages := []*page.PageImpl{}

		for i := 0; i < 20; i++ {
			blankPages = append(blankPages, page.NewPage(i))
		}
		dir := directory.Init(blankPages)

		
		//Scan
		//dir := directory.Directory{BlankPages: blankPages, UsedPages: usedPages}
		dir.Scan()

		// SEEK
		
		content := "d2"
		dir.Seek(content);




}
