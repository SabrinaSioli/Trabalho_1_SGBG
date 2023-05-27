// Trabalho 01 SGBD Prof Zé Maria
// Luiza Clara de Albuquerque Pacheco 493478
// Sabrina Silveira Oliveira 494013

package main

import (
	"Trabalho_1_SGBD/directory"
	"fmt"
)

func main() {
	dir := directory.Init() //Criação do Diretório com as 20 páginas vazias
	var option string
	
	for {
		fmt.Println("\n ******************** MENU ******************** ")
		fmt.Println("[ 1 ] SCAN")
		fmt.Println("[ 2 ] SCAN PAGES")
		fmt.Println("[ 3 ] SEEK")
		fmt.Println("[ 4 ] DELETE")
		fmt.Println("[ 5 ] INSERT")
		fmt.Println("[ 6 ] SAIR")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&option)

		switch option {
		case "1":
		fmt.Println(" ___________________________ SCAN ___________________________ ")
		dir.Scan()
		case "2": 
		fmt.Println(" ___________________________ SCAN PAGES ___________________________ ")
		dir.ScanPages()
		case "3": 
		fmt.Println(" ___________________________ SEEK ___________________________ ")
		fmt.Print("\nConteúdo: ")
		var content string
		fmt.Scan(&content)
		dir.Seek(content)
		case "4":
		fmt.Println(" ___________________________ DELETE ___________________________ ")
		fmt.Print("\nConteúdo: ")
		var content string
		fmt.Scan(&content)
		dir.Delete(content)
		case "5": 
		fmt.Println(" ___________________________ INSERT ___________________________ ")
		var content string
		fmt.Print("Digite o conteúdo: ")
		fmt.Scan(&content)
		if len(content) < 1 || len(content) > 5 {
			fmt.Println("Conteúdo de tamanho inválido.")
		} else {
			dir.Insert(content)
		}
		case "6":
			fmt.Println("Saindo do programa.")
			return
		default:
			fmt.Println("Entrada inválida.")
			option = "0"
		} 
		
	}

	//fmt.Println(" ----------------------- CASOS TESTES ----------------------- ")

	// preencher todas as páginas
	// 1a) inserir 10 "aaa" e 10 "bbb"
	// 3a) scan e scan pages
	// 4a) seek "aaa"
	// 4b) seek "bbb"
	// 4c) seek "aa"

	// completar páginas com espaço
	// 1b) inserir 2 "cc" e 2 "d"
	// 3b) scan e scan pages
	// 4d) seek "cc"
	// 4e) seek "d"
	// 4f) seek "dd"

	// inserir sem espaço suficiente
	// 1c) inserir "eeee"
	// 3c) scan e scan pages
	// 4g) seek "eeee"

	// deletar somente primeira ocorrência + remover das páginas usadas
	// 2a) deletar "bbb"
	// 3d) scan e scan pages
	// 4h) seek "bbb"

	// inserir nas páginas em branco
	// 1d) inserir "fff"
	// 3e) scan e scan pages
	// 4i) seek "fff"

	// atualizar documentos da página após deletar
	// 2b) deletar "d"
	// 3f) scan e scan pages

	// inserir nas páginas usadas
	// 1e) inserir "g"
	// 4j) seek "g"
	// 3g) scan e scan pages

	//fmt.Println(" ----------------------- TESTES SEM MENU ----------------------- ")

	// for i := 0; i < 15; i++ {
	// 	dir.Insert("aaa")
	// }
	// dir.Insert("a")
	// dir.Insert("a")
	// dir.Insert("aaa")
	// dir.Insert("aaa")
	// dir.Insert("aaa")
	// dir.Insert("aaa")
	// dir.Insert("aaa")

	// dir.ScanPages()

	// fmt.Println("Normal")
	// dir.Scan()

	// dir.Seek("aaa")
	// dir.Seek("a")
	// dir.Seek("b")
	// dir.Seek("aaa")

	// dir.Delete("aaa")
	// dir.Delete("a")
	// dir.Delete("aaa")
	// //dir.ScanWithPages()

	// // fmt.Println("Vazias:")
	// // for _, pag := range dir.BlankPages {
	// // 	fmt.Println(pag)
	// // }

	// dir.Insert("aaaaa")
	// dir.ScanPages()

	// fmt.Println("Vazias:")
	// for _, pag := range dir.BlankPages {
	// 	fmt.Println(pag)
	// }
}
