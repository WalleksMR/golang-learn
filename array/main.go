package main

import "fmt"

// A diferenção do array para o slice que o array eu digo qual é o tamanha do array.
func main() {
	list := [4]int{4, 9, 6, 7}

	fmt.Println("List: ", list)

	// # Diferenção do array para o slice
	// -> Depois de ter preenchido o array não é possivel adicionar mais item
	// -> Array ja preenche todas as informação do tamanha que foi defindo
	// -> Array mais performatico com slice

}
