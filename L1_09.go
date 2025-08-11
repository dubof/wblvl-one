package main

import (
	"fmt"
	"log"
	"sync"
)

func generator(wg *sync.WaitGroup, nums []int, out chan<- int) {
	defer wg.Done()
	log.Println("Генератор: начал отправку чисел.")
	for _, n := range nums {
		out <- n
	}
	close(out)
	log.Println("Генератор: закончил отправку, канал закрыт.")
}

func processor(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	log.Println("Обработчик: начал обработку")
	for n := range in {
		result := n * 2
		out <- result
	}
	close(out)
	log.Println("Обработчик: закончил обработку, канал закрыт.")
}

func main() {
	initialData := []int{1, 2, 3, 4, 8, 12}
	numbersCh := make(chan int)
	resultsCh := make(chan int)
	
	var wg sync.WaitGroup
	wg.Add(2)

	go generator(&wg, initialData, numbersCh)
	go processor(&wg, numbersCh, resultsCh)

	log.Println("Main: ожидание результатовв")
	for result := range resultsCh {
		fmt.Printf("Получен результат: %d\n", result)
	}

	wg.Wait()
	log.Println("Main: конвейер завершил работу")
}