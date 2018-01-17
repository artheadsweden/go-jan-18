package main

import (
	"fmt"
	"time"
)

func pinger(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("%v number %v", "Ping", i)
	}
	fmt.Println("Pinger Done")
}

func printer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var ch chan string = make(chan string, 10)

	go pinger(ch)
	go printer(ch)

	time.Sleep(time.Second * 5)
}
