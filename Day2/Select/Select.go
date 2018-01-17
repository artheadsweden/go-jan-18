package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	startValue := 100

	go func() {
		x := startValue
		for {
			c1 <- "From 1: " + strconv.Itoa(x)
			time.Sleep(time.Second * 2)
			x--
		}
	}()

	go func() {
		x := startValue
		for {
			c2 <- "From 2: " + strconv.Itoa(x)
			time.Sleep(time.Millisecond * 500)
			x++
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
			//default:
			//	fmt.Println("Nothing to read")
			//time.Sleep(time.Millisecond * 200)
		}
	}

}
