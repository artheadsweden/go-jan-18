package main

import "fmt"

func goroutine(ch chan string) {
	ch <- "Ping from goroutine 1"
	ch <- "Ping from goroutine 2"
	msg := <-ch
	fmt.Println(msg)
	ch <- "Nice to hear from you"
}

func main() {
	ch := make(chan string)

	//	go func() {
	//		ch <- "Ping from goroutine 1"
	//		ch <- "Ping from goroutine 2"
	//	}()

	go goroutine(ch)
	msg := <-ch
	fmt.Println(msg)
	msg = <-ch
	fmt.Println(msg)
	ch <- "Hi from main"
	msg = <-ch
	fmt.Println(msg)
}
