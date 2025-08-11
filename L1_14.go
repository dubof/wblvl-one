package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Переменная имеет тип int со значением: %d\n", v)
	case string:
		fmt.Printf("Переменная имеет тип string со значением: \"%s\"\n", v)
	case bool:
		fmt.Printf("Переменная имеет тип bool со значением: %t\n", v)
	case chan int:
	
		fmt.Printf("Переменная является каналом типа chan int: %v\n", v)
	case chan string:
		fmt.Printf("Переменная является каналом типа chan string: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T со значением: %v\n", v, v)
	}
}

func main() {

	myInt := 101
	myString := "hello world"
	myBool := true
	myChan := make(chan int)

	checkType(myInt)
	checkType(myString)
	checkType(myBool)
	checkType(myChan)
	checkType(3.14) 
}