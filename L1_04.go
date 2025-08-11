package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ctx context.Context, jobs <-chan int) {
	defer wg.Done() 
	log.Printf("Воркер %d запущен", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Воркер %d: получен сигнал завершения. Выход.", id)
			return
		case job, ok := <-jobs:
			if !ok {
				log.Printf("Воркер %d: канал задач закрыт. Выход.", id)
				return
			}
			log.Printf("Воркер %d: начал обработку задачи %d", id, job)
			time.Sleep(2 * time.Second)
			log.Printf("Воркер %d: закончил обработку задачи %d", id, job)
		}
	}
}

func main() {
	log.Println("Приложение запущено. Нажмите Ctrl+C для завершения.")

	const numWorkers = 3
	const numJobs = 10
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	var wg sync.WaitGroup

	jobs := make(chan int, numJobs)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, ctx, jobs)
	}

	go func() {
		defer close(jobs)

		for i := 1; i <= numJobs; i++ {
			select {
			case jobs <- i:
			case <-ctx.Done():
				log.Println("Генератор задач: получен сигнал завершения. Прекращаем отправку.")
				return
			}
		}
	}()

	<-ctx.Done()

	log.Println("Сигнал завершения получен. Ожидание завершения всех воркеров...")
	wg.Wait()

	log.Println("Все воркеры завершили свою работу. Приложение выключается.")
}