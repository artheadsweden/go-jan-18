package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 6)
	for i := 0; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		t := <-limiter
		fmt.Println("Request", req, time.Now(), " --- ", t)
	}
}
