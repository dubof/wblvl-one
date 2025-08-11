package main

import (
	"fmt"
	"time"
)

func customSleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Привет!")
	customSleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды. Пока!")
}