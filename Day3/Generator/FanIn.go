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

func fanIn(input1, input2 <-chan string) <-chan string {
	outChan := make(chan string)
	go func() {
		for {
			outChan <- <-input1
		}
	}()
	go func() {
		for {
			outChan <- <-input2
		}
	}()
	return outChan
}

func main() {
	out := fanIn(doSomething("Sue"), doSomething("Joe"))
	for i := 0; i < 5; i++ {
		fmt.Printf("Sue: %v \n", <-out)
	}
	fmt.Println("Leaving")
}
