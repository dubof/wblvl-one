package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(len(arr))]

	low := make([]int, 0)
	middle := make([]int, 0)
	high := make([]int, 0)

	for _, item := range arr {
		switch {
		case item < pivot:
			low = append(low, item)
		case item == pivot:
			middle = append(middle, item)
		case item > pivot:
			high = append(high, item)
		}
	}

	low = quickSort(low)
	high = quickSort(high)
	low = append(low, middle...)
	low = append(low, high...)

	return low
}

func main() {
	unsorted := []int{4, 8, 2, 1, 9, 5, 7, 6, 3}
	fmt.Printf("Исходный массив: %v\n", unsorted)

	sorted := quickSort(unsorted)
	fmt.Printf("Отсортированный массив: %v\n", sorted)
}