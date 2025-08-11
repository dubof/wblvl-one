package main

import (
	"context"
	"log"
	"time"
)

func producer(ctx context.Context, dataChan chan<- int) {
	defer close(dataChan)
	
	i := 0
	for {
		time.Sleep(500 * time.Millisecond) 
		
		select {
		case <-ctx.Done():
			log.Println("Продюсер: получил сигнал отмены. Завершаю работу.")
			return
		case dataChan <- i:
			log.Printf("Продюсер: отправил данные -> %d", i)
			i++
		}
	}
}

func main() {
	const N = 5 
	timeout := time.Duration(N) * time.Second
	log.Printf("Программа запустится на %d секунд.", N)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dataChannel := make(chan int)

	go producer(ctx, dataChannel)


	log.Println("Потребитель: начал чтение данных...")

	for data := range dataChannel {
		log.Printf("Потребитель: прочитал данные <- %d", data)
	}

	log.Println("Потребитель: канал закрыт. Программа завершена.")
}