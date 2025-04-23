package main

import (
	"fmt"
	"sort"
)

func main() {
	/*Dado um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1" que vão de 1 a 10, efetuar a soma em duas variaveis, a primeira numeros de 1 a 5 e a segunda de 6 a 10. Imprimir os dois resultados.*/
	numbers := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	// Primeira soma
	upToFive := 0
	upToTen := 0

	for _, number := range numbers {
		if number <= 5 {
			upToFive += number
		} else {
			upToTen += number
		}
	}

	// Ordenando o slice
	sort.Ints(numbers)

	// Imprimindo os resultados
	fmt.Printf("Soma dos primeiros 5 numeros: %d\n", upToFive)
	fmt.Printf("Soma dos ultimos 5 numeros: %d\n", upToTen)
	fmt.Printf("Numeros em ordem crescente: %v\n", numbers)

}
