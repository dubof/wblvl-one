package main

import "fmt"

func removeWithCopy(slice []int, i int) []int {

	if i < 0 || i >= len(slice) {
		return slice
	}
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6}
	indexToRemove := 3 
	fmt.Printf("Исходный слайс: %v\n", mySlice)
	result := removeWithCopy(mySlice, indexToRemove)
	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, result)
}