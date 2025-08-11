package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	mask := int64(1 << pos)
	n |= mask
	return n
}

func clearBit(n int64, pos uint) int64 {
	mask := int64(1 << pos)
	n &^= mask
	return n
}

func main() {
	var number int64 = 5 

	fmt.Printf("Исходное число: %d (%08b)\n", number, number)
	fmt.Println("---------------------------------")

	var posToClear uint = 1
	resultClear := clearBit(number, posToClear)
	fmt.Printf("Сбрасываем бит %d в 0:\n", posToClear)
	fmt.Printf("Результат: %d (%08b)\n", resultClear, resultClear)
	
	fmt.Println("---------------------------------")

	
	var posToSet uint = 2
	resultSet := setBit(number, posToSet)
	fmt.Printf("Устанавливаем бит %d в 1:\n", posToSet)
	fmt.Printf("Результат: %d (%08b)\n", resultSet, resultSet)
	
	fmt.Println("---------------------------------")

	var number2 int64 = 4
	fmt.Printf("Исходное число: %d (%08b)\n", number2, number2)
	resultSet2 := setBit(number2, 0)
	fmt.Printf("Устанавливаем бит 0 в 1:\n")
	fmt.Printf("Результат: %d (%08b)\n", resultSet2, resultSet2)
}