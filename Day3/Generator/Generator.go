package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doSomething(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s => %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func main() {
	ch := doSomething("Do Something")
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine say: %v \n", <-ch)
	}
	fmt.Println("Leaving")
}
