package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greet() { // метод структуры
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct { // наследует
	Human   
	Activity string 
}

func main() {
	a := Action{
		Human: Human{
			Name: "Максим",
			Age:  27,
		},
		Activity: "бегает",
	}


	a.Greet() 	// вызываем метод Greet() напрямую у экземпляра Action

	// прямое обращение
	fmt.Printf("%s сейчас %s.\n", a.Name, a.Activity) 
}

// Привет, меня зовут Максим, мне 27 лет.
// Максим сейчас бегает.