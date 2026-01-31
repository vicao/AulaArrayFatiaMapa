package fatia

import "fmt"

func main() {

	//Primeira parte da aula
	arr := [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}
	fatia := arr[0:5]
	fmt.Println("Fatia:", fatia)
	//Segunda parte da aula
	fatia1 := []int{1, 2, 3}
	fatia2 := append(fatia1, 2, 4, 5)
	fmt.Println("Fatia1:", fatia1, fatia2)
	//Terceira parte da aula
	fatia3 := []int{10, 20, 30, 40, 50}
	fatia4 := make([]int, 1)
	copy(fatia4, fatia3)
	fmt.Println("Fatia3 e Fatia4:", fatia3, fatia4)
}
