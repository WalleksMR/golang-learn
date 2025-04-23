package main

import "fmt"

func main() {
	// Listagem estatica
	list := []int{4, 9, 6, 7}
	fmt.Println("Lista: ", list)
	fmt.Println("Lista pos2: ", list[1])
	fmt.Println("Lista tamanho: ", len(list))

	//Use make quando precisa de um slice de determinado tamanho inicial, mas os valores serão preenchidos depois
	listMake := make([]int, 10)

	// listMake := make([]int, 0) Iniciar com zero elementos

	// Depois preenche os valores
	listMake[0] = 42
	fmt.Println("Lista com make inicada:", len(listMake))
	fmt.Println(listMake[0])

	// Obeter os 3 primeiro items da lista
	fmt.Println("Os 3 primeiros items da lista: ", listMake[:3])

	// Obter os 3 ultimos items da lista
	fmt.Println("Os 3 ultimos items da lista: ", listMake[7:])

	// Adiciona um novo item ao slice
	listMake = append(listMake, 100)
	fmt.Println("Lista com novo item adicionado:", listMake)

}
