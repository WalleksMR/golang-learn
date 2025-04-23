package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 5}
	slice2 := []string{"abacaxi", "abobora", "banana", "uva"}

	fmt.Println(slice1)
	fmt.Println(reverse(slice1))
	fmt.Printf("\n")
	fmt.Println(slice2)
	fmt.Println(reverse(slice2))
}

func reverse[T int | string](slice []T) []T {
	newSlice := make([]T, len(slice))

	newSliceLen := len(newSlice) - 1

	for _, value := range slice {
		newSlice[newSliceLen] = value
		newSliceLen--
	}

	return newSlice
}
