// 1. Закрытие канала (использование for-range)

package main

import (
	"fmt"
	"log"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	log.Printf("Воркер %d: запущен", id)
	for job := range jobs {
		fmt.Printf("Воркер %d: обработал задачу %d\n", id, job)
	}
	log.Printf("Воркер %d: канал закрыт, завершаю работу.", id)
}

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(1, &wg, jobs)

	for i := 1; i <= 3; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
	log.Println("Программа завершена.")
}

// 2. Канал для уведомления (Done/Quit Channel)
package main

import (
	"log"
	"time"
)

func worker(done <-chan struct{}) {
	log.Println("Воркер: запущен")
	for {
		select {
		case <-done:
			log.Println("Воркер: получил сигнал, завершаю работу.")
			return
		default:
			log.Println("Воркер: работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker(done)
	time.Sleep(2 * time.Second)

	log.Println("main: отправляю сигнал завершения.")
	close(done)

	time.Sleep(1 * time.Second)
	log.Println("Программа завершена.")
}

// 3. Использование context.Context
package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Воркер: запущен")
	for {
		select {
		case <-ctx.Done():
			log.Println("Воркер: контекст отменен, завершаю работу.")
			return
		default:
			log.Println("Воркер: работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	time.Sleep(2 * time.Second)

	log.Println("main: отменяю контекст.")
	cancel()

	wg.Wait()
	log.Println("Программа завершена.")
}

// 4. runtime.Goexit() - редкий и специфичный
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	defer fmt.Println("Воркер: defer выполнен!")

	fmt.Println("Воркер: начинаю работу.")
	fmt.Println("Воркер: что-то пошло не так, нужно выйти")
	
	runtime.Goexit()
	fmt.Println("Воркер: эта строка не будет напечатана")
}

func main() {
	go worker()
	
	time.Sleep(1 * time.Second) 
	fmt.Println("main: программа завершена.")
}

// 5. 5. os.Exit() - радикальное прерывание
package main

import (
	"log"
	"os"
	"time"
)

func worker() {
	defer log.Println("Воркер: defer НЕ будет выполнен!")
	
	log.Println("Воркер: работаю...")
	time.Sleep(1 * time.Second)
	
	log.Println("Воркер: обнаружена фатальная ошибка!")
	os.Exit(1) 
}

func main() {
	log.Println("main: запуск воркера.")
	go worker()
	select {}
}