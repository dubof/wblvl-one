package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {

		mid := low + (high-low)/2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			low = mid + 1
		case arr[mid] > target:
			high = mid - 1
		}
	}

	return -1
}

func main() {
	sortedSlice := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	target1 := 23
	index1 := binarySearch(sortedSlice, target1)
	fmt.Printf(" %d в %v -> Индекс: %d\n", target1, sortedSlice, index1)

	target2 := 91
	index2 := binarySearch(sortedSlice, target2)
	fmt.Printf(" %d в %v -> Индекс: %d\n", target2, sortedSlice, index2)

	target3 := 100
	index3 := binarySearch(sortedSlice, target3)
	fmt.Printf(" %d в %v -> Индекс: %d\n", target3, sortedSlice, index3)
}