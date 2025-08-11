package main

import "fmt"

func intersection(a, b []int) []int {
	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	resultSet := make(map[int]struct{})
	for _, v := range b {
		if _, ok := set[v]; ok {
			resultSet[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(resultSet))
	for v := range resultSet {
		result = append(result, v)
	}

	return result
}

func main() {
	sliceA := []int{1, 2, 3, 5, 8}
	sliceB := []int{2, 3, 4, 8, 9, 10}

	result := intersection(sliceA, sliceB)

	fmt.Printf("Слайс A: %v\n", sliceA)
	fmt.Printf("Слайс B: %v\n", sliceB)
	fmt.Printf("Пересечение: %v\n", result)
}