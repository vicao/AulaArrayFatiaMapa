package mapa

import "fmt"

func Mapa() {
	//Primeira parte da aula
	x := make(map[string]int)
	x["Chave"] = 10
	fmt.Println(x["Chave"])
	//Segunda parte da aula
	y := make(map[int]int)
	y[1] = 20
	fmt.Println(y[1])
	//Terceira parte da aula
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["Li"])
}

