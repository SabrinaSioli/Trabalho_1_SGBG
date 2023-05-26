package main

import (
	"Trabalho_1_SGBD/directory"
	"Trabalho_1_SGBD/page"
	"fmt"
)

const MaxPage = 20

func main() {
	//Testes Iniciais
	/*
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

	*/
	// Insert Pages
	blankPages := []*page.PageImpl{}

	for i := 0; i < 20; i++ {
		blankPages = append(blankPages, page.NewPage(i))
	}
	dir := directory.Init(blankPages)

	for i := 0; i < 15; i++ {
		dir.Insert("aaa")
	}
	dir.Insert("a")
	dir.Insert("a")
	dir.Insert("aaa")
	dir.Insert("aaa")
	dir.Insert("aaa")
	dir.Insert("aaa")
	dir.Insert("aaa")

	dir.ScanPages()

	fmt.Println("Normal")
	dir.Scan()

	fmt.Println("Seek")
	dir.Seek("aaa")
	dir.Seek("a")
	dir.Seek("b")
	dir.Seek("aaa")

	dir.Delete("aaa")
	dir.Delete("a")
	dir.Delete("aaa")
	//dir.ScanWithPages()

	// fmt.Println("Vazias:")
	// for _, pag := range dir.BlankPages {
	// 	fmt.Println(pag)
	// }

	dir.Insert("aaaaa")
	dir.ScanPages()

	fmt.Println("Vazias:")
	for _, pag := range dir.BlankPages {
		fmt.Println(pag)
	}



	/*
	//Scan
	//dir := directory.Directory{BlankPages: blankPages, UsedPages: usedPages}
	dir.Scan()

	// INSERT
	fmt.Println("******************* INSERT *******************")

	for i := 0; i < 5; i++ {
		dir.Insert("aaa")
	}
	dir.Insert("x1")
	dir.Insert("x2")
	dir.Insert("bbb1")
	dir.Insert("bbb2")
	dir.Insert("bbb3")
	dir.Insert("bbb4")
	dir.Insert("bbb5")
	dir.Insert("bbb6")

	for i := 0; i < 9; i++ {
		dir.Insert("aaa")
	}
	
	for i := 0; i < 5; i++ {
		dir.Delete("aaa")
	}

	dir.Delete("x1")
	dir.Delete("x2")
	dir.Delete("bbb1")
	dir.Delete("bbb2")
	dir.Delete("bbb3")
	dir.Delete("bbb4")
	dir.Delete("bbb5")
	dir.Delete("bbb6")

	fmt.Println("Vazias:")
	for _, pag := range dir.BlankPages {
		fmt.Println(pag)
	}

	fmt.Println("Usadas:")
	for _, pag := range dir.UsedPages {
		fmt.Println(pag)
	}

	/*
	dir.Insert("iii1")
	dir.Insert("iii2")

	dir.ScanWithPages()

	fmt.Println("Vazias:")
	for _, pag := range dir.BlankPages {
		fmt.Println(pag)
	}

	fmt.Println("Usadas:")
	for _, pag := range dir.UsedPages {
		fmt.Println(pag)
	}
	fmt.Printf("")
	dir.ScanWithPages()
*/
}
