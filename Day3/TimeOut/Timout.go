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
			time.Sleep(time.Duration(rand.Intn(1100)) * time.Millisecond)
		}
	}()
	return ch
}

func main() {
	ch := doSomething("Joe")
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(time.Second * 1):
			fmt.Println("So slow, I'm out")
			return
		}
	}
}
