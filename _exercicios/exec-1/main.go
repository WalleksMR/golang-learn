package main

import "fmt"

func main() {
	numbers := [2]int32{1, 2}
	var total int32
	for _, n := range numbers {
		total += n
	}
	fmt.Println("Total:", total)
}
