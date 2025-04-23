package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Não possoui o termo while, mas o equivalente em go segue o exmeplo abaixo
	text := "palavra"
	total := len(text)
	index := 0

	for index < total {
		fmt.Println("index: ", index)
		index++
	}
}
