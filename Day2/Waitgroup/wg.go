package main

import (
	"fmt"
	"sync"
	"time"
)

func SleepFunc(name string, sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine", name, "starting...")
	time.Sleep(sec * time.Second)
	fmt.Println("Goroutine", name, "exit")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go SleepFunc("First", 1, &wg)
	go SleepFunc("Second", 2, &wg)

	wg.Wait()
	fmt.Println("Main goroutine exit")
}
