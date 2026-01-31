package main

import (
	"AulaArrayFatiaMapa/array"
	"AulaArrayFatiaMapa/fatia"
	"AulaArrayFatiaMapa/mapa"
	"fmt"
)

func ShowMenu() {
	fmt.Println("=== Menu Principal ===")
	fmt.Println("1. Código Array")
	fmt.Println("2. Código Fatia/Slice")
	fmt.Println("3. Código Map")
	fmt.Println("5. Sair")

}

func main() {
	for {
		ShowMenu()
		fmt.Print("Escolha uma opção: ")
		var opcao int
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Entrada inválida, tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			array.Array()
		case 2:
			fatia.Fatia()
		case 3:
			mapa.Mapa()
		case 5:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
