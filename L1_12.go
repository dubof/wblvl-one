package main

import "fmt"

func main() {
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}

    set := make(map[string]struct{})
    for _, item := range sequence {
        set[item] = struct{}{}
    }

    result := make([]string, 0, len(set))
    for key := range set {
        result = append(result, key)
    }

    fmt.Printf("Исходная последовательность: %v\n", sequence)
    fmt.Printf("Полученное множество: %v\n", result) 
}