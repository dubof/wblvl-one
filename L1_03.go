package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ctx context.Context, dataStream <-chan int) {
	defer wg.Done()
	log.Printf("Воркер %d запущен", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Воркер %d: получен сигнал завершения. Выход.", id)
			return
		case data := <-dataStream:
			fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: Не указано количество воркеров.")
		fmt.Println("Использование: go run . <количество_воркеров>")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Printf("Ошибка: Некорректное количество воркеров: %s\n", os.Args[1])
		os.Exit(1)
	}

	log.Printf("Запускаем %d воркеров. Нажмите Ctrl+C для завершения.", numWorkers)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel() 

	var wg sync.WaitGroup

	dataChannel := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, ctx, dataChannel)
	}

	go func() {
		defer close(dataChannel)
		
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				log.Println("Продюсер: получен сигнал завершения. Прекращаем отправку данных.")
				return
			case dataChannel <- i:

				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()

	log.Println("Сигнал завершения получен. Ожидание завершения всех воркеров...")
	wg.Wait()
	log.Println("Все воркеры завершили свою работу. Приложение выключается.")
}

//go run .\L1_03.go 10