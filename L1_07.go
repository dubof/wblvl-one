package main

import (
	"fmt"
	"log"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex 
	v  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		v: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()         
	defer sm.mu.Unlock() 
	sm.v[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()         
	defer sm.mu.Unlock() 
	val, ok := sm.v[key]
	return val, ok
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup
	const numWriters = 1000 

	log.Println("Запуск горутин для записи в безопасную карту")
	wg.Add(numWriters)

	for i := 0; i < numWriters; i++ {
		go func(val int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", val)
			safeMap.Set(key, val)
		}(i)
	}

	wg.Wait()
	log.Printf("Запись завершена. Карта содержит %d элементов.", len(safeMap.v))

	val, ok := safeMap.Get("key-50")
	if ok {
		log.Printf("Значение для 'key-50': %d", val)
	}
}

// go run .\L1_07.go  -race . 