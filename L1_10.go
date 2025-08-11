package main

import (
"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, t := range temperatures {
		key := int(t/10) * 10
		groups[key] = append(groups[key], t)
	}

	fmt.Println("Результат группировки температур:")
	fmt.Println(groups)
}