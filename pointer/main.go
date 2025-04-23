package main

import "fmt"

/*
Sobre os ponteiros
*/
func main() {
	fmt.Printf("========Main\n")
	x := 5
	y := x
	y = 10

	fmt.Println("Valores: ", "x:", x, " ", "y:", y)
	fmt.Println("Referencia na memoria de uma variavel", "x:", &x, "y:", &y)
	fmt.Printf("\n")

	/*
		Quando uma variavel receber o valor de outra variavel, elas serão independentes, pois elas fazer referencia a alocação de memorias diferentes
		E para saber qual referencia da memoria de uma variavel, basta utiliar o e comercial "&"
	*/

	fmt.Println("Referencia na memoria de uma variavel", "x:", &x, "y:", &y)
	fmt.Printf("\n")

	/*
	   Porem temos um problema quando usando funções, pois ele usar outro espaço na memoria
	*/
	SumOne(x, y)

	/*
	   Para usar na mesma referencia da memoria, precisamos usar o asterisco *
	*/
	newY := &x
	fmt.Println("Referencia na memoria de uma variavel newY", "x:", &x, "newY:", &*newY)
	SumTwo(&x, newY)
}

func SumOne(x int, y int) {
	fmt.Printf("========SumOne\n")
	fmt.Println("Valores: ", "x:", x, " ", "y:", y)
	fmt.Printf("\n")

	fmt.Println("Referencia na memoria de uma variavel", "x:", &x, "y:", &y)
	fmt.Printf("\n")
}

func SumTwo(x *int, y *int) {
	fmt.Printf("========SumTwo\n")
	fmt.Println("Valores: ", "x:", *x, " ", "y:", *y)
	fmt.Printf("\n")

	fmt.Println("Referencia na memoria de uma variavel", "x:", &*x, "y:", &*y)
	fmt.Printf("\n")
}
