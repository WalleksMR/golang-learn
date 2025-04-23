package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x interface{} = 42
	reflectTypeOf(x)

	//---------------------------------------------

	typeAssertions(42)
	typeAssertions("texto")

	//---------------------------------------------

	fmtPackage()

	//---------------------------------------------

}

/*
1. reflect.TypeOf() (Pacote reflect)

			Características:
	   Retorna um valor do tipo reflect.Type (não uma string)
	   Útil para inspeção dinâmica de tipos
	   Requer cuidado com tipos de interface (interface{})
*/
func reflectTypeOf(x interface{}) {
	fmt.Println("reflectTypeOf", reflect.TypeOf(x))

	y := "Hello"
	fmt.Println("reflectTypeOf", reflect.TypeOf(y))
}

/*
 2. Type Assertions (Asserções de Tipo)
    Articles:
    https://medium.com/@jamal.kaksouri/mastering-type-assertion-in-go-a-comprehensive-guide-216864b4ea4d
*/
func typeAssertions(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("typeAssertions", "É um inteiro")
	case string:
		fmt.Println("typeAssertions", "É uma string")
	default:
		fmt.Println("typeAssertions", "Tipo desconhecido")
	}
}

// 3. %T no fmt Package
func fmtPackage() {
	x := 3.14
	fmt.Printf("fmtPackage Tipo: %T\n", x) // Saída: Tipo: float64
}
