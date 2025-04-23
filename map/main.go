package main

import "fmt"

// Map é um objeto de chave e valor
// Onde declatore qual  tipo da chave ([string]) e o valor que é inteiro (int)
func main() {
	// Uma forma de criar o map coms os valores
	fmt.Printf("Uma forma de criar o map coms os valores\n\n")

	populations := map[string]int{"PA": 40000, "RJ": 78009878, "MG": 899000}

	fmt.Println("Map é um objeto de chave e valor. populations: ", populations)

	fmt.Println("Para acessar a população do PA basta acessar neste formato: \" populations[\"PA\"]\" populations PA: ", populations["PA"])
	fmt.Printf("\n")
	// Uma forma de criar o map quando não sabe os valores
	fmt.Printf("Uma forma de criar o map quando não sabe os valores\n\n")
	populations2 := make(map[string]int)
	populations2["SP"] = 3000000
	populations2["DF"] = 100000

	fmt.Println("Map vazio, populations2: ", populations2)

	// Para adicionar novos valores ao map
	populations2["RS"] = 1000000
	fmt.Println("Map com novo item populations2: ", populations2)
	fmt.Printf("\n")
	// E para percorrer pelo map, basta usar o range
	fmt.Printf("E para percorrer pelo map, basta usar o range\n\n")

	for state, population := range populations2 {
		fmt.Printf("A população de %s é %d\n", state, population)
	}
	fmt.Printf("\n")

	// E para remover um item populations2
	// passo a variavel e depois a chave que desenja remover
	delete(populations2, "DF")
	fmt.Println("Map após remover DF populations2: ", populations2)
	fmt.Printf("\n")
	// Para verificar se um estado existe no map

}
