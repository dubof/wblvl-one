package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex 
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()         
	defer c.mu.Unlock()
	c.value++
}


func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	numGoroutines := 1000 

	wg.Add(numGoroutines)


	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счётчика (Mutex): %d\n", counter.Value())
}