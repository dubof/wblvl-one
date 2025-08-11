package main

import "fmt"

func main() {
    a, b := 10, 20
    fmt.Printf("Исходные значения: a = %d, b = %d\n", a, b)

    a = a + b // 30 (10 + 20)
    b = a - b // 10 (30 - 20)
    a = a - b // 20 (30 - 10)

    fmt.Printf("Новые значения:    a = %d, b = %d\n", a, b)
}