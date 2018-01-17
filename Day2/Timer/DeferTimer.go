package main

import (
	"fmt"
	"time"
)

func startTimer() func() {
	fmt.Println("Starting timer")
	t := time.Now()

	return func() {
		fmt.Println("It took", time.Now().Sub(t), "seconds")
	}
}

func main() {
	x := "Hepp"
	func() {
		defer startTimer()()
		fmt.Println(x)
		time.Sleep(time.Second * 3)
	}()
	fmt.Println("This is not timed")
}
